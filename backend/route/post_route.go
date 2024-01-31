package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupPostRoute(app fiber.Router, postController *controller.PostController) {
	app.Post("/posts", postController.Create)
	app.Get("/posts/:code", postController.Get)
	app.Put("/posts/:code", postController.Update)
}
