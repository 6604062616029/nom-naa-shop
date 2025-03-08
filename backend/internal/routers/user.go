package routers

import (
	"github.com/baimhons/nom-naa-shop.git/internal/handlers"
	"github.com/baimhons/nom-naa-shop.git/internal/middlewares"
	"github.com/baimhons/nom-naa-shop.git/internal/validations"
	"github.com/gofiber/fiber/v2"
)

type UserRountes struct {
	app            fiber.Router
	userHandler    *handlers.UserHandler
	validate       *validations.UserValidateImpl
	authMiddleware *middlewares.AuthMiddleware
	userValidate   *validations.UserValidateImpl
}

func NewUserRountes(app fiber.Router, userHandler *handlers.UserHandler, validate *validations.UserValidateImpl, authMiddleware *middlewares.AuthMiddleware) *UserRountes {
	return &UserRountes{app: app, userHandler: userHandler, validate: validate, authMiddleware: authMiddleware}
}

func (r *UserRountes) SetupRoutes() {
	users := r.app.Group("/users")
	users.Post("/register", r.validate.ValidateRegisterUser, r.userHandler.RegisterUser)
	users.Post("/login", r.validate.ValidateLoginUser, r.userHandler.LoginUser)
	users.Post("/logout", r.authMiddleware.AuthToken, r.userHandler.LogoutUser)
	users.Get("/profile", r.authMiddleware.AuthToken, r.userHandler.GetUserProfile)
	users.Get("/all", r.authMiddleware.AuthToken, r.userValidate.ValidateRoleAdmin, r.userHandler.GetAllUsers)
	users.Put("/profile", r.authMiddleware.AuthToken, r.validate.ValidateUpdateUser, r.userHandler.UpdateUser)
	users.Delete("/:id", r.authMiddleware.AuthToken, r.validate.ValidateRoleAdmin, r.userHandler.DeleteUser)
}
