package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/kimnopal/maya/repository"
	"gorm.io/gorm"
)

type MajorService struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	MajorRepository *repository.MajorRepository
}

func (s *MajorService) Create() {

}
