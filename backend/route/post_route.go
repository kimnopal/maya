package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupPostRoute(app fiber.Router, postController *controller.PostController) {
	app.Post("/posts", postController.Create)
}
