package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type PostCategoryRepository struct {
	Repository[entity.PostCategory]
}

func (r *PostCategoryRepository) FindByName(DB *gorm.DB, postCategory *entity.PostCategory, name string) error {
	return DB.Where("name = ?", name).Take(postCategory).Error
}
