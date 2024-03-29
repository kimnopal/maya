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

type PostService struct {
	DB                     *gorm.DB
	Validate               *validator.Validate
	PostRepository         *repository.PostRepository
	PostCategoryRepository *repository.PostCategoryRepository
}

func NewPostService(DB *gorm.DB, Validate *validator.Validate, PostRepository *repository.PostRepository, PostCategoryRepository *repository.PostCategoryRepository) *PostService {
	return &PostService{
		DB:                     DB,
		Validate:               Validate,
		PostRepository:         PostRepository,
		PostCategoryRepository: PostCategoryRepository,
	}
}

func (s *PostService) Create(ctx context.Context, request *model.PostCreateRequest) (*model.PostResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	postCategory := new(entity.PostCategory)
	if err := s.PostCategoryRepository.FindById(tx, postCategory, request.PostCategoryID); err != nil {
		return nil, fiber.ErrNotFound
	}

	post := converter.PostCreateRequestToEntity(request)
	if err := s.PostRepository.Create(tx, post); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostEntityToResponse(post), nil
}

func (s *PostService) Get(ctx context.Context, request *model.PostGetRequest) (*model.PostResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	post := new(entity.Post)
	if err := s.PostRepository.FindByCode(tx, post, request.Code); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostEntityToResponse(post), nil
}

func (s *PostService) Update(ctx context.Context, request *model.PostUpdateRequest) (*model.PostResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	postCategory := new(entity.PostCategory)
	if err := s.PostCategoryRepository.FindById(tx, postCategory, request.PostCategoryID); err != nil {
		return nil, fiber.ErrNotFound
	}

	post := new(entity.Post)
	if err := s.PostRepository.FindByCode(tx, post, request.Code); err != nil {
		return nil, fiber.ErrNotFound
	}

	post = converter.PostUpdateRequestToEntity(post, request)
	if err := s.PostRepository.Update(tx, post); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostEntityToResponse(post), nil
}

func (s *PostService) Delete(ctx context.Context, request *model.PostDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return fiber.ErrBadRequest
	}

	post := new(entity.Post)
	if err := s.PostRepository.FindByCode(tx, post, request.Code); err != nil {
		return fiber.ErrNotFound
	}

	if err := s.PostRepository.Delete(tx, post); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
