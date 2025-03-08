package handlers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.RegisterUser)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	user, statusCode, err := h.userService.RegisterUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User registered successfully",
		Data:    user,
	})
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.LoginUser)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	user, statusCode, err := h.userService.LoginUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User logged in successfully",
		Data:    user,
	})
}

func (h *UserHandler) LogoutUser(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)

	statusCode, err := h.userService.LogoutUser(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User logged out successfully",
	})
}

func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)

	user, statusCode, err := h.userService.GetUserProfile(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User profile fetched successfully",
		Data:    user,
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, statusCode, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "Users fetched successfully",
		Data:    users,
	})
}
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.UpdateUser)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	user, statusCode, err := h.userService.UpdateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User updated successfully",
		Data:    user,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	statusCode, err := h.userService.DeleteUser(uuid.MustParse(id))
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "User deleted successfully",
	})
}
