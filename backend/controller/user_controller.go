package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(UserService *service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.UserRegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	userResponse, err := c.UserService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", userResponse))
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	userResponse, err := c.UserService.Login(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", userResponse))
}
