package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
)

func SetupFacultyRoute(app fiber.Router, facultyController *controller.FacultyController) {
	app.Post("/faculties", facultyController.Create)
	app.Put("/faculties/:id", facultyController.Update)
	app.Delete("/faculties/:id", facultyController.Delete)
	app.Get("/faculties/:id", facultyController.Get)
	app.Get("/faculties", facultyController.GetAll)
}
