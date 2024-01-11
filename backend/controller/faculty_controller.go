package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/service"
)

type FacultyController struct {
	FacultyService *service.FacultyService
}

func NewFacultyController(FacultyService *service.FacultyService) *FacultyController {
	return &FacultyController{
		FacultyService: FacultyService,
	}
}

func (c *FacultyController) Create(ctx *fiber.Ctx) error {
	request := new(model.FacultyCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON(converter.ToWebResponse(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, nil))
	}

	facultyResponse, err := c.FacultyService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.JSON(converter.ToWebResponse(fiber.StatusOK, "OK", facultyResponse))
}

func (c *FacultyController) Update(ctx *fiber.Ctx) error {
	request := new(model.FacultyUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}
	request.ID = uint64(id)

	facultyResponse, err := c.FacultyService.Update(ctx.UserContext(), request)
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.JSON(converter.ToWebResponse(fiber.StatusOK, "OK", facultyResponse))
}

func (c *FacultyController) Delete(ctx *fiber.Ctx) error {
	request := new(model.FacultyDeleteRequest)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}

	request.ID = uint64(id)
	if err := c.FacultyService.Delete(ctx.UserContext(), request); err != nil {
		return ctx.JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.JSON(converter.ToWebResponse(fiber.StatusOK, "OK", nil))
}

func (c *FacultyController) Get(ctx *fiber.Ctx) error {
	request := new(model.FacultyGetRequest)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}

	request.ID = uint64(id)
	facultyResponse, err := c.FacultyService.Get(ctx.UserContext(), request)
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.JSON(converter.ToWebResponse(fiber.StatusOK, "OK", facultyResponse))
}

func (c *FacultyController) GetAll(ctx *fiber.Ctx) error {
	facultiesResponse, err := c.FacultyService.GetAll(ctx.UserContext())
	if err != nil {
		return ctx.JSON(converter.ToWebResponse(err.(*fiber.Error).Code, err.(*fiber.Error).Message, nil))
	}

	return ctx.JSON(converter.ToWebResponse(fiber.StatusOK, "OK", facultiesResponse))
}

// belum selesai
func (c *FacultyController) ListWithMajors(ctx *fiber.Ctx) error {
	faculties, err := c.FacultyService.ListWithMajors(ctx.UserContext())
	if err != nil {
		return err
	}
	return ctx.JSON(converter.ToWebResponse(200, "OK", faculties))
}
