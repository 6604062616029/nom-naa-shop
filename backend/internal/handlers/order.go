package handlers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) ConfirmOrder(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)

	req := c.Locals("req").(request.OrderRequest)

	order := models.Order{
		AddressID:     req.AddressID,
		CartID:        req.CartID,
		PaymentMethod: req.PaymentMethod,
	}

	order, status, err := h.orderService.ConfirmOrder(order, userContext)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order confirmed successfully",
		"order":   order,
	})
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {

	req := c.Locals("req").(request.UpdateOrderStatusRequest)

	order := models.Order{
		BaseModel: models.BaseModel{
			ID: req.OrderID,
		},
		Status: req.Status,
	}

	order, status, err := h.orderService.UpdateOrderStatus(order)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order status updated successfully",
		"order":   order,
	})
}

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	order, status, err := h.orderService.GetOrderByID(uuid.MustParse(orderID))
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order fetched successfully",
		"order":   order,
	})
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	querys := c.Locals("querys").(request.PaginationQuery)

	response, status, err := h.orderService.GetAllOrders(querys)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *OrderHandler) GetTotalRevenue(c *fiber.Ctx) error {
	totalRevenue, status, err := h.orderService.GetTotalRevenue()
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Total revenue fetched successfully",
		"totalRevenue": totalRevenue,
	})
}
