package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func RoleEntityToResponse(role *entity.Role) *model.RoleResponse {
	return &model.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

func RoleCreateRequestToEntity(createRequest *model.RoleCreateRequest) *entity.Role {
	return &entity.Role{
		Name: createRequest.Name,
	}
}

func RoleUpdateRequestToEntity(role *entity.Role, updateRequest *model.RoleUpdateRequest) *entity.Role {
	role.Name = updateRequest.Name

	return role
}
