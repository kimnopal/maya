package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupContactTypeRoute(app fiber.Router, contactTypeController *controller.ContactTypeController) {
	app.Post("/contact-types", contactTypeController.Create)
}
