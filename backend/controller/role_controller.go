package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type RoleController struct {
	RoleService *service.RoleService
}

func NewRoleController(RoleService *service.RoleService) *RoleController {
	return &RoleController{
		RoleService: RoleService,
	}
}

func (c *RoleController) Create(ctx *fiber.Ctx) error {
	request := new(model.RoleCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	roleResponse, err := c.RoleService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", roleResponse))
}
