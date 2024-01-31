package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	Repository[entity.Post]
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) FindByCode(DB *gorm.DB, post *entity.Post, code string) error {
	return DB.Where("code = ?", code).Take(post).Error
}
