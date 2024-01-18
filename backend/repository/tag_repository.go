package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type TagRepository struct {
	Repository[entity.Tag]
}

func NewTagRepository() *TagRepository {
	return &TagRepository{}
}

func (r *TagRepository) FindByName(DB *gorm.DB, tag *entity.Tag, name string) error {
	return DB.Where("name = ?", name).Take(tag).Error
}
