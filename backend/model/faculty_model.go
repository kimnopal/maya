package model

import "github.com/kimnopal/maya/entity"

type FacultyResponse struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	CreatedAt uint64          `json:"created_at"`
	UpdatedAt uint64          `json:"updated_at"`
	Majors    []*entity.Major `json:"majors,omitempty"`
}

type FacultyCreateRequest struct {
	Name string `json:"name" validate:"required,max:100"`
}

type FacultyUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max:100"`
}

type FacultyDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type FacultyGetRequest struct {
	ID uint64 `json:"id" validate:"required"`
}
