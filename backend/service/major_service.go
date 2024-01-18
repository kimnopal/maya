package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/repository"
	"gorm.io/gorm"
)

type MajorService struct {
	DB                *gorm.DB
	Validate          *validator.Validate
	FacultyRepository *repository.FacultyRepository
	MajorRepository   *repository.MajorRepository
}

func NewMajorService(DB *gorm.DB, Validate *validator.Validate, FacultyRepository *repository.FacultyRepository, MajorRepository *repository.MajorRepository) *MajorService {
	return &MajorService{
		DB:                DB,
		Validate:          Validate,
		FacultyRepository: FacultyRepository,
		MajorRepository:   MajorRepository,
	}
}

func (s *MajorService) Create(ctx context.Context, request *model.MajorCreateRequest) (*model.MajorResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	faculty := new(entity.Faculty)
	if err := s.FacultyRepository.FindById(tx, faculty, request.FacultyID); err != nil {
		return nil, fiber.ErrNotFound
	}

	major := converter.MajorCreateRequestToEntity(request)
	if err := s.MajorRepository.Create(tx, major); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.MajorEntityToResponse(major), nil
}

func (s *MajorService) Get(ctx context.Context, request *model.MajorGetRequest) (*model.MajorResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	major := new(entity.Major)
	if err := s.MajorRepository.FindByIdWithFaculty(tx, major, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}
	fmt.Println(major)

	return converter.MajorEntityToResponse(major), nil
}

func (s *MajorService) Delete(ctx context.Context, request *model.MajorDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return fiber.ErrBadRequest
	}

	major := new(entity.Major)
	if err := s.MajorRepository.FindById(tx, major, request.ID); err != nil {
		return fiber.ErrNotFound
	}

	if err := s.MajorRepository.Delete(tx, major); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *MajorService) Update(ctx context.Context, request *model.MajorUpdateRequest) (*model.MajorResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	faculty := new(entity.Faculty)
	if err := c.FacultyRepository.FindById(tx, faculty, request.FacultyID); err != nil {
		return nil, fiber.ErrNotFound
	}

	major := new(entity.Major)
	if err := c.MajorRepository.FindById(tx, major, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	major = converter.MajorUpdateRequestToEntity(major, request)
	if err := c.MajorRepository.Update(tx, major); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.MajorEntityToResponse(major), nil
}

func (s *MajorService) GetAll(ctx context.Context) ([]*model.MajorResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	majors := new([]*entity.Major)
	if err := s.MajorRepository.FindAll(tx, majors); err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.MajorListEntityToResponse(majors), nil
}
