package service

import "gorm.io/gorm"

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Create() {

}
