package validations

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderValidationImpl struct {
}

type OrderValidation interface {
	ValidateConfirmOrder(c *fiber.Ctx) error
	ValidateUpdateOrderStatus(c *fiber.Ctx) error
}

func NewOrderValidation() OrderValidation {
	return &OrderValidationImpl{}
}

func (v *OrderValidationImpl) ValidateConfirmOrder(c *fiber.Ctx) error {
	var req request.OrderRequest

	validateCommonRequestBody(c, &req)

	if req.CartID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Cart ID is required",
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (v *OrderValidationImpl) ValidateUpdateOrderStatus(c *fiber.Ctx) error {
	var req request.UpdateOrderStatusRequest

	validateCommonRequestBody(c, &req)

	c.Locals("req", req)
	return c.Next()
}
