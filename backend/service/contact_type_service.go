package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/repository"
	"gorm.io/gorm"
)

type ContactTypeService struct {
	DB                    *gorm.DB
	Validate              *validator.Validate
	ContactTypeRepository *repository.ContactTypeRepository
}

func NewContactTypeService(DB *gorm.DB, Validate *validator.Validate, ContactTypeRepository *repository.ContactTypeRepository) *ContactTypeService {
	return &ContactTypeService{
		DB:                    DB,
		Validate:              Validate,
		ContactTypeRepository: ContactTypeRepository,
	}
}

func (s *ContactTypeService) Create(ctx context.Context, request *model.ContactTypeCreateRequest) (*model.ContactTypeResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	contactType := converter.ContactTypeCreateRequestToEntity(request)
	if err := s.ContactTypeRepository.Create(tx, contactType); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactTypeEntityToResponse(contactType), nil
}
