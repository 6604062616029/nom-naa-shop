package handlers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

func (h *CartHandler) AddItemToCart(c *fiber.Ctx) error {

	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	req, ok := c.Locals("req").(request.AddItemToCartRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	resp, statusCode, err := h.cartService.AddItemToCart(req, userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Item added to cart successfully",
		Data:    resp,
	})
}

func (h *CartHandler) GetCartByID(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	resp, statusCode, err := h.cartService.GetCartByID(uuid.MustParse(userContext.ID))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Cart fetched successfully",
		Data:    resp,
	})
}

func (h *CartHandler) UpdateItemFromCart(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	req, ok := c.Locals("req").(request.UpdateItemFromCartRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	resp, statusCode, err := h.cartService.UpdateItemFromCart(req, userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Item updated from cart successfully",
		Data:    resp,
	})
}

func (h *CartHandler) DeleteItemFromCart(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	itemID := c.Params("item_id")

	resp, statusCode, err := h.cartService.DeleteItemFromCart(uuid.MustParse(itemID), userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Item deleted from cart successfully",
		Data:    resp,
	})
}

func (h *CartHandler) ConfirmCart(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	req, ok := c.Locals("req").(request.ConfirmCartRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	resp, statusCode, err := h.cartService.ConfirmCart(req.CartID, userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Cart confirmed successfully",
		Data:    resp,
	})
}

func (h *CartHandler) CancelCart(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized",
		})
	}

	resp, statusCode, err := h.cartService.CancelCart(uuid.MustParse(userContext.ID), userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Cart cancelled successfully",
		Data:    resp,
	})
}
