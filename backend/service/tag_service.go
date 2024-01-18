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

type TagService struct {
	DB            *gorm.DB
	Validate      *validator.Validate
	TagRepository *repository.TagRepository
}

func NewTagService(DB *gorm.DB, Validate *validator.Validate, TagRepository *repository.TagRepository) *TagService {
	return &TagService{
		DB:            DB,
		Validate:      Validate,
		TagRepository: TagRepository,
	}
}

func (s *TagService) Create(ctx context.Context, request *model.TagCreateRequest) (*model.TagResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	tag := converter.TagCreateRequestToEntity(request)
	if err := s.TagRepository.Create(tx, tag); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagEntityToResponse(tag), nil
}

func (s *TagService) Update(ctx context.Context, request *model.TagUpdateRequest) (*model.TagResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := s.TagRepository.FindById(tx, tag, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	tag = converter.TagUpdateRequestToEntity(tag, request)
	if err := s.TagRepository.Update(tx, tag); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagEntityToResponse(tag), nil
}

func (s *TagService) Delete(ctx context.Context, request *model.TagDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := s.TagRepository.FindById(tx, tag, request.ID); err != nil {
		return fiber.ErrNotFound
	}

	if err := s.TagRepository.Delete(tx, tag); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *TagService) Get(ctx context.Context, request *model.TagGetRequest) (*model.TagResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := s.TagRepository.FindByName(tx, tag, request.Name); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagEntityToResponse(tag), nil
}

func (s *TagService) GetAll(ctx context.Context) ([]*model.TagResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	tags := new([]*entity.Tag)
	if err := s.TagRepository.FindAll(tx, tags); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagListEntityToResponse(tags), nil
}
