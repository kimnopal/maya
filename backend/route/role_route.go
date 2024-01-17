package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupRoleRoute(app fiber.Router, roleController *controller.RoleController) {
	app.Get("/roles", roleController.Create)
	app.Put("/roles/:id", roleController.Update)
	app.Delete("/roles/:id", roleController.Delete)
}
