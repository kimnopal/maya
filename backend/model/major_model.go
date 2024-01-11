package model

import "github.com/kimnopal/maya/entity"

type MajorResponse struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt uint64         `json:"created_at"`
	UpdatedAt uint64         `json:"updated_at"`
	Faculty   entity.Faculty `json:"faculty,omitempty"`
}

type MajorCreateRequest struct {
	Name string `json:"name" validate:"required,max:100"`
}

type MajorUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max:100"`
}

type MajorDeleteRequest struct {
	ID uint64 `json:"id"`
}

type MajorGetRequest struct {
	ID uint64 `json:"id"`
}
