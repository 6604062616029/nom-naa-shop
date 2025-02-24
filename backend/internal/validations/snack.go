package validations

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/gofiber/fiber/v2"
)

type SnackValidate interface {
	ValidateCreateSnackRequest(c *fiber.Ctx) error
	ValidateUpdateSnackRequest(c *fiber.Ctx) error
}

type SnackValidateImpl struct{}

func NewSnackValidate() *SnackValidateImpl {
	return &SnackValidateImpl{}
}

var AllowedImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
}

func (v *SnackValidateImpl) ValidateCreateSnackRequest(c *fiber.Ctx) error {
	var req request.CreateSnackRequest

	validateCommonRequestFormBody(c, &req)

	if err := validateImageFiles(req.Files); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (v *SnackValidateImpl) ValidateUpdateSnackRequest(c *fiber.Ctx) error {
	var req request.UpdateSnackRequest

	validateCommonRequestFormBody(c, &req)

	if err := validateImageFiles(req.Files); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
