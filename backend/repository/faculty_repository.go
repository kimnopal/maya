package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type FacultyRepository struct {
	Repository[entity.Faculty]
}

func NewFacultyRepository() *FacultyRepository {
	return &FacultyRepository{}
}

// belum selesai
func (r *FacultyRepository) FindAllWithMajors(db *gorm.DB, faculties *[]entity.Faculty) error {
	return db.Preload("Majors").Find(faculties).Error
}

func (r *FacultyRepository) FindByIdWithMajors(db *gorm.DB, faculty *entity.Faculty, id any) error {
	return db.Preload("Majors").Where("id = ?", id).Take(faculty).Error
}
