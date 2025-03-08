package validations

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/gofiber/fiber/v2"
)

type PaymentValidate interface {
	ValidatePayment(c *fiber.Ctx) error
}

type PaymentValidateImpl struct{}

func NewPaymentValidate() PaymentValidate {
	return &PaymentValidateImpl{}
}

func (v *PaymentValidateImpl) ValidatePayment(c *fiber.Ctx) error {
	var req request.PaymentRequest

	validateCommonRequestFormBody(c, &req)

	if len(req.Proof) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "at least one proof file is required",
		})
	}

	if err := validateImageFiles(req.Proof); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
