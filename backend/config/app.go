package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kimnopal/maya/controller"
	"github.com/kimnopal/maya/middleware"
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
	roleRepository := repository.NewRoleRepository()
	userRepository := repository.NewUserRepository()
	tagRepository := repository.NewTagRepository()
	postCategoryRepository := repository.NewPostCategoryRepository()
	postRepository := repository.NewPostRepository()

	facultyService := service.NewFacultyService(config.DB, config.Validate, facultyRepository)
	majorService := service.NewMajorService(config.DB, config.Validate, facultyRepository, majorRepository)
	roleService := service.NewRoleService(config.DB, config.Validate, roleRepository)
	userService := service.NewUserService(config.DB, config.Validate, userRepository)
	tagService := service.NewTagService(config.DB, config.Validate, tagRepository)
	postCategoryService := service.NewPostCategoryService(config.DB, config.Validate, postCategoryRepository)
	postService := service.NewPostService(config.DB, config.Validate, postRepository, postCategoryRepository)

	facultyController := controller.NewFacultyController(facultyService)
	majorController := controller.NewMajorController(majorService)
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService)
	tagController := controller.NewTagController(tagService)
	postCategoryController := controller.NewPostCategoryController(postCategoryService)
	postController := controller.NewPostController(postService)

	api := config.App.Group("/api")
	corsConfig := cors.ConfigDefault
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = "http://localhost:3000"
	api.Use(cors.New(corsConfig))
	route.SetupUserRoute(api, userController)

	api.Use(middleware.AuthMiddleware)
	route.SetupFacultyRoute(api, facultyController)
	route.SetupMajorRoute(api, majorController)
	route.SetupRoleRoute(api, roleController)
	route.SetupTagRoute(api, tagController)
	route.SetupPostCategoryRoute(api, postCategoryController)
	route.SetupPostRoute(api, postController)
}
