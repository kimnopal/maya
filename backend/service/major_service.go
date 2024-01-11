package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/repository"
	"gorm.io/gorm"
)

type MajorService struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	MajorRepository *repository.MajorRepository
}

func (s *MajorService) Create(ctx context.Context, request *model.MajorCreateRequest) (*model.MajorResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	major := new(entity.Major)
	if err := s.MajorRepository.Create(tx, major); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.MajorEntityToResponse(major), nil
}
