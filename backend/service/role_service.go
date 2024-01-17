package service

import (
	"github.com/go-playground/validator/v10"
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

// func (s *RoleService)Create(ctx context.Context, request)  {

// }
