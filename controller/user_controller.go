package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/model"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) *UserController {
	return &UserController{
		DB: DB,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.UserRegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.JSON(nil)
	}

	return nil
}
