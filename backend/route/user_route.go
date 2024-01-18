package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
	"github.com/kimnopal/maya/middleware"
)

func SetupUserRoute(app fiber.Router, userController *controller.UserController) {
	app.Get("/users/register", userController.Register)
	app.Post("/users/login", userController.Login)

	app.Use(middleware.AuthMiddleware)
}
