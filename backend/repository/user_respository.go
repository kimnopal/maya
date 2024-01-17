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

func (r *UserRepository) FindByUsername(DB *gorm.DB, user *entity.User) error {
	return DB.Where("username = ?", user.Username).Take(user).Error
}
