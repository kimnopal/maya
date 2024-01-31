package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/helper"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type PostController struct {
	PostService *service.PostService
}

func NewPostController(PostService *service.PostService) *PostController {
	return &PostController{
		PostService: PostService,
	}
}

func (c *PostController) Create(ctx *fiber.Ctx) error {
	request := new(model.PostCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	request.UserID = ctx.Locals("auth").(model.UserResponse).ID
	request.Code = helper.GenerateRandomCode(20)

	postResponse, err := c.PostService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", postResponse))
}

func (c *PostController) Get(ctx *fiber.Ctx) error {
	request := new(model.PostGetRequest)
	request.Code = ctx.Params("code")

	postResponse, err := c.PostService.Get(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", postResponse))
}

func (c *PostController) Update(ctx *fiber.Ctx) error {
	request := new(model.PostUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}
	request.Code = ctx.Params("code")

	postResponse, err := c.PostService.Update(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(err.(*fiber.Error).Code).JSON(converter.ToWebResponse(
			err.(*fiber.Error).Code,
			err.(*fiber.Error).Message,
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(converter.ToWebResponse(fiber.StatusOK, "OK", postResponse))
}
