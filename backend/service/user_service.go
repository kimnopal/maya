package service

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	if err := s.UserRespository.FindByUsername(tx, user, request.Username); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fiber.ErrNotFound
	}

	if user.ID != 0 {
		return nil, fiber.ErrBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	request.Password = string(hashedPassword)
	user = converter.UserRegisterRequestToEntity(request)
	user.RoleID = 1

	if err := s.UserRespository.Create(tx, user); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserEntityToResponse(user), nil
}

func (s *UserService) Login(ctx context.Context, request *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := s.UserRespository.FindByUsername(tx, user, request.Username); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	claims := new(model.UserClaims)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 48))

	claims.UserResponse = *converter.UserEntityToResponse(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	finalToken, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserLoginResponse{Token: finalToken}, nil
}
