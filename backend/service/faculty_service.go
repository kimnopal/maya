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

type FacultyService struct {
	DB                *gorm.DB
	Validate          *validator.Validate
	FacultyRepository *repository.FacultyRepository
}

func NewFacultyService(DB *gorm.DB, Validate *validator.Validate, FacultyRepository *repository.FacultyRepository) *FacultyService {
	return &FacultyService{
		DB:                DB,
		Validate:          Validate,
		FacultyRepository: FacultyRepository,
	}
}

func (s *FacultyService) Create(ctx context.Context, request *model.FacultyCreateRequest) (*model.FacultyResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	faculty := converter.FacultyCreateRequestToEntity(request)
	if err := s.FacultyRepository.Create(tx, faculty); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.FacultyEntityToResponse(faculty), nil
}

func (s *FacultyService) Update(ctx context.Context, request *model.FacultyUpdateRequest) (*model.FacultyResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	faculty := new(entity.Faculty)

	if err := s.FacultyRepository.FindById(tx, faculty, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	faculty = converter.FacultyUpdateRequestToEntity(faculty, request)

	if err := s.FacultyRepository.Update(tx, faculty); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.FacultyEntityToResponse(faculty), nil
}

func (s *FacultyService) Delete(ctx context.Context, request *model.FacultyDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return fiber.ErrBadRequest
	}

	faculty := new(entity.Faculty)
	if err := s.FacultyRepository.FindById(tx, faculty, request.ID); err != nil {
		return fiber.ErrNotFound
	}

	if err := s.FacultyRepository.Delete(tx, faculty); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *FacultyService) Get(ctx context.Context, request *model.FacultyGetRequest) (*model.FacultyResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	faculty := new(entity.Faculty)
	if err := s.FacultyRepository.FindById(tx, faculty, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.FacultyEntityToResponse(faculty), nil
}

func (s *FacultyService) GetAll(ctx context.Context) ([]*model.FacultyResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	faculties := new([]*entity.Faculty)
	if err := s.FacultyRepository.FindAll(tx, faculties); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.FacultyListEntityToResponse(faculties), nil
}

// belum selesai
func (s *FacultyService) ListWithMajors(ctx context.Context) (*[]entity.Faculty, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var faculties []entity.Faculty
	if err := s.FacultyRepository.FindAllWithMajors(tx, &faculties); err != nil {
		return nil, fiber.ErrNotFound
	}

	// var facultiesResponse []*model.FacultyResponse
	return &faculties, nil
}
