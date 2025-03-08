package handlers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	req := c.Locals("req").(request.PaymentRequest)

	payment, err := req.ToModel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process payment proof: " + err.Error(),
		})
	}

	result, err := h.paymentService.CreatePayment(payment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Payment created successfully",
		"payment": result,
	})
}

func (h *PaymentHandler) GetPaymentProof(c *fiber.Ctx) error {
	id := c.Params("id")
	payment, err := h.paymentService.GetPaymentByID(uuid.MustParse(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Payment not found",
		})
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Send(payment.Proof)
}
