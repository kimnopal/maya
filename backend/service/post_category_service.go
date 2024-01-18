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

type PostCategoryService struct {
	DB                     *gorm.DB
	Validate               *validator.Validate
	PostCategoryRepository *repository.PostCategoryRepository
}

func NewPostCategoryService(DB *gorm.DB, Validate *validator.Validate, PostCategoryRepository *repository.PostCategoryRepository) *PostCategoryService {
	return &PostCategoryService{
		DB:                     DB,
		Validate:               Validate,
		PostCategoryRepository: PostCategoryRepository,
	}
}

func (s *PostCategoryService) Create(ctx context.Context, request *model.PostCategoryCreateRequest) (*model.PostCategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	postCategory := converter.PostCategoryCreateRequestToEntity(request)
	if err := s.PostCategoryRepository.Create(tx, postCategory); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostCategoryEntityToResponse(postCategory), nil
}

func (s *PostCategoryService) Update(ctx context.Context, request *model.PostCategoryUpdateRequest) (*model.PostCategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	postCategory := new(entity.PostCategory)
	if err := s.PostCategoryRepository.FindById(tx, postCategory, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	postCategory = converter.PostCategoryUpdateRequestToEntity(postCategory, request)
	if err := s.PostCategoryRepository.Update(tx, postCategory); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostCategoryEntityToResponse(postCategory), nil
}

func (s *PostCategoryService) Delete(ctx context.Context, request *model.PostCategoryDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return fiber.ErrBadRequest
	}

	postCategory := new(entity.PostCategory)
	if err := s.PostCategoryRepository.FindById(tx, postCategory, request.ID); err != nil {
		return fiber.ErrNotFound
	}

	if err := s.PostCategoryRepository.Delete(tx, postCategory); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PostCategoryService) Get(ctx context.Context, request *model.PostCategoryGetRequest) (*model.PostCategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	postCategory := new(entity.PostCategory)
	if err := s.PostCategoryRepository.FindByName(tx, postCategory, request.Name); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostCategoryEntityToResponse(postCategory), nil
}
