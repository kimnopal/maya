package repository

import "github.com/kimnopal/maya/entity"

type RoleRepository struct {
	Repository[entity.Role]
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}
