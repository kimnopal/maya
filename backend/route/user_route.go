package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupUserRoute(app fiber.Router, userController *controller.UserController) {
	app.Get("/users/register", userController.Register)
}
