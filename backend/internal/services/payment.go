package services

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentService interface {
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetPaymentByID(id uuid.UUID) (*models.Payment, error)
}

type PaymentServiceImpl struct {
	paymentRepository repositories.PaymentRepository
}

func NewPaymentService(paymentRepository repositories.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{paymentRepository: paymentRepository}
}

func (s *PaymentServiceImpl) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	if err := s.paymentRepository.Create(payment); err != nil {
		return nil, err
	}

	// Fetch the created payment without the nested order data
	var createdPayment models.Payment
	if err := s.paymentRepository.GetDB().
		Select("id, create_at, update_at, delete_at, order_id, amount, proof").
		First(&createdPayment, payment.ID).Error; err != nil {
		return nil, err
	}

	return &createdPayment, nil
}

func (s *PaymentServiceImpl) GetPaymentByID(id uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	if err := s.paymentRepository.GetDB().
		Preload("Order", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, create_at, update_at, delete_at, tracking_id, cart_id, total_price, status, payment_method, address_id").Omit("Payment")
		}).
		First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
