package service

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}
