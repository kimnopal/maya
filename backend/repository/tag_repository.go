package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type TagRepository struct {
	Repository[entity.Tag]
}

func (r *TagRepository) FindByValue(DB *gorm.DB, tag *entity.Tag, value string) error {
	return DB.Where("value = ?", value).Take(tag).Error
}
