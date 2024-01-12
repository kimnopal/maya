package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupMajorRoute(app fiber.Router, majorController *controller.MajorController) {
	app.Post("/majors", majorController.Create)
}
