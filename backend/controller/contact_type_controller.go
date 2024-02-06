package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type ContactTypeController struct {
	ContactTypeService *service.ContactTypeService
}

func NewContactTypeController(ContactTypeService *service.ContactTypeService) *ContactTypeController {
	return &ContactTypeController{
		ContactTypeService: ContactTypeService,
	}
}

func (c *ContactTypeController) Create(ctx *fiber.Ctx) error {
	request := new(model.ContactTypeCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	contactTypeResponse, err := c.ContactTypeService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", contactTypeResponse))
}
