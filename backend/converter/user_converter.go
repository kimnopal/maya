package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func UserEntityToResponse(entity *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:       entity.ID,
		Username: entity.Username,
		Name:     entity.Name,
		Email:    entity.Email,
	}
}

func UserRegisterRequestToEntity(registerRequest *model.UserRegisterRequest) *entity.User {
	return &entity.User{
		Username: registerRequest.Username,
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}
