package repository

import "github.com/kimnopal/maya/entity"

type UserRepository struct {
	Repository[entity.User]
}
