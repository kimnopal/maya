package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type TagController struct {
	TagService *service.TagService
}

func NewTagController(TagService *service.TagService) *TagController {
	return &TagController{
		TagService: TagService,
	}
}

func (c *TagController) Create(ctx *fiber.Ctx) error {
	request := new(model.TagCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	tagResponse, err := c.TagService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", tagResponse))
}

func (c *TagController) Update(ctx *fiber.Ctx) error {
	request := new(model.TagUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	request.ID = uint64(id)

	tagResponse, err := c.TagService.Update(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", tagResponse))
}

func (c *TagController) Delete(ctx *fiber.Ctx) error {
	request := new(model.TagDeleteRequest)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	request.ID = uint64(id)

	err = c.TagService.Delete(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", nil))
}

func (c *TagController) Get(ctx *fiber.Ctx) error {
	request := new(model.TagGetRequest)
	name := ctx.Params("name")
	request.Name = name

	tagResponse, err := c.TagService.Get(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", tagResponse))
}

func (c *TagController) GetAll(ctx *fiber.Ctx) error {
	tagResponses, err := c.TagService.GetAll(ctx.UserContext())
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", tagResponses))
}
