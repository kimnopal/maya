package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupPostCategoryRoute(app fiber.Router, postCategoryController *controller.PostCategoryController) {
	app.Post("/post-categories", postCategoryController.Create)
	app.Get("/post-categories/:name", postCategoryController.Get)
	app.Put("/post-categories/:id", postCategoryController.Update)
	app.Delete("/post-categories/:id", postCategoryController.Delete)
}
