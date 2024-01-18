package model

type TagResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type TagCreateRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type TagUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100"`
}

type TagDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type TagGetRequest struct {
	Name string `json:"id" validate:"required"`
}
