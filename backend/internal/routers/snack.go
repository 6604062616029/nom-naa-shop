package routers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/handlers"
	"github.com/baimhons/nom-naa-shop.git/internal/middlewares"
	"github.com/baimhons/nom-naa-shop.git/internal/validations"
	"github.com/gofiber/fiber/v2"
)

type SnackRouter struct {
	app            fiber.Router
	snackHandler   *handlers.SnackHandler
	authMiddleware *middlewares.AuthMiddleware
	userValidate   *validations.UserValidateImpl
	snackValidate  *validations.SnackValidateImpl
}

func NewSnackRouter(app fiber.Router, snackHandler *handlers.SnackHandler, authMiddleware *middlewares.AuthMiddleware, userValidate *validations.UserValidateImpl, snackValidate *validations.SnackValidateImpl) *SnackRouter {
	return &SnackRouter{app: app, snackHandler: snackHandler, authMiddleware: authMiddleware, userValidate: userValidate, snackValidate: snackValidate}
}

func (r *SnackRouter) SetupRoutes() {
	snack := r.app.Group("/snack")
	snack.Post("/", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, validations.NewSnackValidate().ValidateCreateSnackRequest, r.snackHandler.CreateSnack)
	snack.Get("/", r.snackHandler.GetAllSnacks)
	snack.Get("/image/:id", r.snackHandler.GetSnackImage)
	snack.Put("/:id", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, validations.NewSnackValidate().ValidateUpdateSnackRequest, r.snackHandler.UpdateSnack)
	snack.Delete("/:id", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, r.snackHandler.DeleteSnack)
	snack.Post("/review", r.authMiddleware.AuthToken, r.snackValidate.ValidateCreateReviewRequest, r.snackHandler.CreateReview)
	snack.Get("/review/:snack_id", r.snackValidate.ValidateGetAllReviewsRequest, r.snackHandler.GetAllReviews)
}
