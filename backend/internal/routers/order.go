package routers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/handlers"
	"github.com/baimhons/nom-naa-shop.git/internal/middlewares"
	"github.com/baimhons/nom-naa-shop.git/internal/validations"
	"github.com/gofiber/fiber/v2"
)

type OrderRouter struct {
	app            fiber.Router
	orderHandler   *handlers.OrderHandler
	authMiddleware *middlewares.AuthMiddleware
	userValidate   *validations.UserValidateImpl
}

func NewOrderRouter(app fiber.Router, orderHandler *handlers.OrderHandler, authMiddleware *middlewares.AuthMiddleware, userValidate *validations.UserValidateImpl) *OrderRouter {
	return &OrderRouter{app: app, orderHandler: orderHandler, authMiddleware: authMiddleware, userValidate: userValidate}
}

func (r *OrderRouter) SetupRoutes() {
	order := r.app.Group("/order")
	order.Post("/confirm", r.authMiddleware.AuthToken, validations.NewOrderValidation().ValidateConfirmOrder, r.orderHandler.ConfirmOrder)
	order.Put("/status", r.authMiddleware.AuthToken, validations.NewOrderValidation().ValidateUpdateOrderStatus, r.orderHandler.UpdateOrderStatus)
	order.Get("/total-revenue", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, r.orderHandler.GetTotalRevenue)
	order.Get("/", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, r.userValidate.ValidateGetUsersRequest, r.orderHandler.GetAllOrders)
	order.Get("/:id", r.authMiddleware.AuthToken, r.orderHandler.GetOrder)
}
