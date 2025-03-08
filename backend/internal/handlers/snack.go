package handlers

import (
	"strconv"

	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SnackHandler struct {
	snackService services.SnackService
}

func NewSnackHandler(snackService services.SnackService) *SnackHandler {
	return &SnackHandler{snackService: snackService}
}

func (h *SnackHandler) CreateSnack(c *fiber.Ctx) error {
	req := c.Locals("req").(request.CreateSnackRequest)
	resp, statusCode, err := h.snackService.CreateSnack(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Failed to create snack: " + err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}

func (h *SnackHandler) GetSnackImage(c *fiber.Ctx) error {
	id := c.Params("id")
	snack, statusCode, err := h.snackService.GetSnackByID(uuid.MustParse(id))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Snack not found",
		})
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Send(snack.Data.(models.Snack).Image)
}

func (h *SnackHandler) GetAllSnacks(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid page parameter",
		})
	}
	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid page_size parameter",
		})
	}
	sort := c.Query("sort", "name")
	order := c.Query("order", "asc")

	querys := request.PaginationQuery{
		Page:     &page,
		PageSize: &pageSize,
		Sort:     &sort,
		Order:    &order,
	}

	resp, statusCode, err := h.snackService.GetAllSnacks(querys)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}

func (h *SnackHandler) UpdateSnack(c *fiber.Ctx) error {
	id := c.Params("id")
	req := c.Locals("req").(request.UpdateSnackRequest)
	userContext := c.Locals("userContext").(models.UserContext)
	resp, statusCode, err := h.snackService.UpdateSnack(req, userContext, uuid.MustParse(id))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Failed to update snack: " + err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}

func (h *SnackHandler) DeleteSnack(c *fiber.Ctx) error {
	id := c.Params("id")
	resp, statusCode, err := h.snackService.DeleteSnack(uuid.MustParse(id))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Failed to delete snack: " + err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}

func (h *SnackHandler) CreateReview(c *fiber.Ctx) error {
	req := c.Locals("req").(request.CreateReviewRequest)
	userContext := c.Locals("userContext").(models.UserContext)
	resp, statusCode, err := h.snackService.CreateReview(req, userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Failed to create review: " + err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}

func (h *SnackHandler) GetAllReviews(c *fiber.Ctx) error {
	querys := c.Locals("querys").(request.PaginationQuery)
	snackID := c.Params("snack_id")
	resp, statusCode, err := h.snackService.GetAllReviewsBySnackID(querys, uuid.MustParse(snackID))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: "Failed to get all reviews: " + err.Error(),
		})
	}
	return c.Status(statusCode).JSON(resp)
}
