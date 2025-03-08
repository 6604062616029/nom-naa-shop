package services

import (
	"errors"
	"net/http"

	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderService interface {
	ConfirmOrder(order models.Order, userContext models.UserContext) (models.Order, int, error)
	UpdateOrderStatus(order models.Order) (models.Order, int, error)
	GetOrderByID(orderID uuid.UUID) (models.Order, int, error)
	GetAllOrders(querys request.PaginationQuery) (response.SuccessResponse, int, error)
	GetTotalRevenue() (float64, int, error)
}

type OrderServiceImpl struct {
	orderRepository repositories.OrderRepository
	cartRepository  repositories.CartRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, cartRepository repositories.CartRepository) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		cartRepository:  cartRepository,
	}
}

func (s *OrderServiceImpl) ConfirmOrder(order models.Order, userContext models.UserContext) (models.Order, int, error) {
	cart, err := s.cartRepository.GetCartByCondition("id = ? AND status = ?", order.CartID, "confirmed")
	if err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	if cart == nil || cart.ID == uuid.Nil {
		return models.Order{}, http.StatusBadRequest, errors.New("cart not found")
	}

	if cart.Status != "confirmed" {
		return models.Order{}, http.StatusBadRequest, errors.New("cart is not confirmed")
	}

	cartWithItems, err := s.cartRepository.GetCartWithItems(cart.ID)
	if err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}
	cart = cartWithItems

	tx := s.cartRepository.Begin()

	var totalPrice float64
	for _, item := range cart.Items {
		totalPrice += item.Snack.Price * float64(item.Quantity)

		if err := tx.Model(&models.Snack{}).
			Where("id = ?", item.SnackID).
			Update("quantity", tx.Raw("quantity - ?", item.Quantity)).
			Error; err != nil {
			tx.Rollback()
			return models.Order{}, http.StatusInternalServerError, err
		}
	}

	if err := tx.Model(&models.Cart{}).Where("id = ?", cart.ID).Update("status", "ordered").Error; err != nil {
		tx.Rollback()
		return models.Order{}, http.StatusInternalServerError, err
	}

	order.CartID = cart.ID
	order.TrackingID = uuid.New().String()
	order.TotalPrice = totalPrice
	order.Status = "pending"

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return models.Order{}, http.StatusInternalServerError, err
	}

	userUUID, err := uuid.Parse(userContext.ID)
	if err != nil {
		tx.Rollback()
		return models.Order{}, http.StatusInternalServerError, err
	}

	if err := tx.Create(&models.Cart{
		UserID: userUUID,
		Status: "pending",
	}).Error; err != nil {
		tx.Rollback()
		return models.Order{}, http.StatusInternalServerError, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	var finalOrder models.Order
	if err := s.orderRepository.GetDB().
		Preload("Cart").
		Preload("Cart.Items").
		Preload("Cart.Items.Snack").
		Preload("Cart.User").
		Preload("Address").
		First(&finalOrder, order.ID).Error; err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	return finalOrder, http.StatusOK, nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(order models.Order) (models.Order, int, error) {
	// First, get the existing order from the database
	var existingOrder models.Order
	if err := s.orderRepository.GetByID(&existingOrder, order.ID); err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	// Update only the status field
	existingOrder.Status = order.Status

	if err := s.orderRepository.Update(&existingOrder); err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	return existingOrder, http.StatusOK, nil
}

func (s *OrderServiceImpl) GetOrderByID(orderID uuid.UUID) (models.Order, int, error) {
	var order models.Order
	if err := s.orderRepository.GetByID(&order, orderID); err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	// Preload payment but select only specific fields to avoid circular reference
	if err := s.orderRepository.GetDB().
		Preload("Payment", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, create_at, update_at, delete_at, order_id, amount, proof").Omit("Order")
		}).
		First(&order, orderID).Error; err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}

	return order, http.StatusOK, nil
}

func (s *OrderServiceImpl) GetAllOrders(querys request.PaginationQuery) (response.SuccessResponse, int, error) {
	var orders []models.Order
	query := s.orderRepository.GetDB().
		Preload("Cart").
		Preload("Cart.Items").
		Preload("Cart.Items.Snack").
		Preload("Cart.User").
		Preload("Address").
		Preload("Payment", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, create_at, update_at, delete_at, order_id, amount, proof").Omit("Order")
		})

	if querys.Page != nil && querys.PageSize != nil {
		offset := *querys.Page * *querys.PageSize
		query = query.Offset(offset).Limit(*querys.PageSize)
	}

	if querys.Sort != nil && querys.Order != nil {
		orderClause := *querys.Sort + " " + *querys.Order
		query = query.Order(orderClause)
	}

	if err := query.Find(&orders).Error; err != nil {
		return response.SuccessResponse{}, http.StatusInternalServerError, err
	}

	return response.SuccessResponse{
		Message: "Orders fetched successfully",
		Data:    orders,
	}, http.StatusOK, nil
}

func (s *OrderServiceImpl) GetTotalRevenue() (float64, int, error) {
	var totalRevenue float64
	if err := s.orderRepository.GetDB().Model(&models.Order{}).Select("SUM(total_price) as total_revenue").Scan(&totalRevenue).Error; err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return totalRevenue, http.StatusOK, nil
}
