package routers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/handlers"
	"github.com/baimhons/nom-naa-shop.git/internal/middlewares"
	"github.com/baimhons/nom-naa-shop.git/internal/validations"
	"github.com/gofiber/fiber/v2"
)

type PaymentRouter struct {
	app            fiber.Router
	paymentHandler *handlers.PaymentHandler
	authMiddleware *middlewares.AuthMiddleware
}

func NewPaymentRouter(app fiber.Router, paymentHandler *handlers.PaymentHandler, authMiddleware *middlewares.AuthMiddleware) *PaymentRouter {
	return &PaymentRouter{app: app, paymentHandler: paymentHandler, authMiddleware: authMiddleware}
}

func (r *PaymentRouter) SetupRoutes() {
	payment := r.app.Group("/payment")
	payment.Post("/create", r.authMiddleware.AuthToken, validations.NewPaymentValidate().ValidatePayment, r.paymentHandler.CreatePayment)
	payment.Get("/proof/:id", r.authMiddleware.AuthToken, r.paymentHandler.GetPaymentProof)
}
