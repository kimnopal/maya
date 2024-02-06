package model

type ContactTypeResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ContactTypeCreateRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type ContactTypeUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type ContactTypeDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type ContactTypeGetRequest struct {
	ID uint64 `json:"id" validate:"required"`
}
