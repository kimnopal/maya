package repository

import (
	"github.com/kimnopal/maya/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByUsername(DB *gorm.DB, user *entity.User, username string) error {
	return DB.Where("username = ?", username).Take(user).Error
}
