package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type MajorController struct {
	MajorService *service.MajorService
}

func NewMajorController(MajorService *service.MajorService) *MajorController {
	return &MajorController{
		MajorService: MajorService,
	}
}

func (c *MajorController) Create(ctx *fiber.Ctx) error {
	request := new(model.MajorCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		fmt.Println("parsing error")
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}

	majorResponse, err := c.MajorService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", majorResponse))
}

func (c *MajorController) Get(ctx *fiber.Ctx) error {
	request := new(model.MajorGetRequest)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}
	request.ID = uint64(id)

	majorResponse, err := c.MajorService.Get(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", majorResponse))
}

func (c *MajorController) Delete(ctx *fiber.Ctx) error {
	request := new(model.MajorDeleteRequest)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}
	request.ID = uint64(id)

	if err := c.MajorService.Delete(ctx.UserContext(), request); err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", nil))
}
