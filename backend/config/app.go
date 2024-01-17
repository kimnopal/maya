package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/controller"
	"github.com/kimnopal/maya/repository"
	"github.com/kimnopal/maya/route"
	"github.com/kimnopal/maya/service"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	facultyRepository := repository.NewFacultyRepository()
	majorRepository := repository.NewMajorRepository()
	userRepository := repository.NewUserRepository()

	facultyService := service.NewFacultyService(config.DB, config.Validate, facultyRepository)
	majorService := service.NewMajorService(config.DB, config.Validate, facultyRepository, majorRepository)
	userService := service.NewUserService(config.DB, config.Validate, userRepository)

	facultyController := controller.NewFacultyController(facultyService)
	majorController := controller.NewMajorController(majorService)
	userController := controller.NewUserController(userService)

	api := config.App.Group("/api")
	route.SetupFacultyRoute(api, facultyController)
	route.SetupMajorRoute(api, majorController)
	route.SetupUserRoute(api, userController)
}
