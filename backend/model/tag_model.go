package model

type TagResponse struct {
	ID    uint64 `json:"id"`
	Value string `json:"value"`
}

type TagCreateRequest struct {
	Value string `json:"value" validate:"required,max=100"`
}

type TagUpdateRequest struct {
	ID    uint64 `json:"id" validate:"required"`
	Value string `json:"value" validate:"required,max=100"`
}

type TagDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type TagGetRequest struct {
	Value string `json:"id" validate:"required"`
}
