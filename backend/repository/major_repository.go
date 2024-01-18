package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type MajorRepository struct {
	Repository[entity.Major]
}

func NewMajorRepository() *MajorRepository {
	return &MajorRepository{}
}

func (r *MajorRepository) FindByIdWithFaculty(DB *gorm.DB, major *entity.Major, id any) error {
	return DB.Preload("Faculty").Where("id = ?", id).Take(major).Error
}
