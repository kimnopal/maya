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

type RoleService struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	RoleRepository *repository.RoleRepository
}

func NewRoleService(DB *gorm.DB, Validate *validator.Validate, RoleRepository *repository.RoleRepository) *RoleService {
	return &RoleService{
		DB:             DB,
		Validate:       Validate,
		RoleRepository: RoleRepository,
	}
}

func (s *RoleService) Create(ctx context.Context, request *model.RoleCreateRequest) (*model.RoleResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	role := converter.RoleCreateRequestToEntity(request)
	if err := s.RoleRepository.Create(tx, role); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleEntityToResponse(role), nil
}

func (s *RoleService) Update(ctx context.Context, request *model.RoleUpdateRequest) (*model.RoleResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	role := new(entity.Role)
	if err := s.RoleRepository.FindById(tx, role, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	role = converter.RoleUpdateRequestToEntity(role, request)
	if err := s.RoleRepository.Update(tx, role); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleEntityToResponse(role), nil
}
