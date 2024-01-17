package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
	"github.com/kimnopal/maya/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	UserRespository *repository.UserRepository
}

func NewUserService(DB *gorm.DB, Validate *validator.Validate, UserRespository *repository.UserRepository) *UserService {
	return &UserService{
		DB:              DB,
		Validate:        Validate,
		UserRespository: UserRespository,
	}
}

func (s *UserService) Create(ctx context.Context, request *model.UserRegisterRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	user.Username = request.Username
	if err := s.UserRespository.FindByUsername(tx, user); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("error find by username")
		return nil, fiber.ErrInternalServerError
	}

	if user.ID != 0 {
		return nil, fiber.ErrBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error hashing password")
		return nil, fiber.ErrInternalServerError
	}

	request.Password = string(hashedPassword)
	user = converter.UserRegisterRequestToEntity(request)
	// user.RoleID = 1

	if err := s.UserRespository.Create(tx, user); err != nil {
		fmt.Println("error create user")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserEntityToResponse(user), nil
}
