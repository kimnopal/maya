package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupTagRoute(app fiber.Router, tagController *controller.TagController) {
	app.Get("/tags", tagController.GetAll)
	app.Get("/tags/:name", tagController.Get)
	app.Post("/tags", tagController.Create)
	app.Delete("/tags/:id", tagController.Delete)
	app.Put("/tags/:id", tagController.Update)
}
