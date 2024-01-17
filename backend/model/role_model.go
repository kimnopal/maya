package model

type RoleResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type RoleCreateRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type RoleUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100"`
}

type RoleDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type RoleGetRequest struct {
	ID uint64 `json:"id" validate:"required"`
}
