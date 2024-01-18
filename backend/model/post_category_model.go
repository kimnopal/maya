package model

type PostCategoryResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type PostCategoryCreateRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type PostCategoryUpdateRequest struct {
	ID   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100"`
}

type PostCategoryDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type PostCategoryGetRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}
