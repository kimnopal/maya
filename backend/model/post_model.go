package model

type PostResponse struct {
	ID          uint64        `json:"id"`
	Title       string        `json:"title"`
	Code        string        `json:"code"`
	Description string        `json:"description"`
	CreatedAt   uint64        `json:"created_at"`
	UpdatedAt   uint64        `json:"updated_at"`
	User        *UserResponse `json:"user"`
}

type PostCreateRequest struct {
	Title          string `json:"title" validate:"required,max=255"`
	Code           string `json:"code" validate:"required,max=100"`
	Description    string `json:"description" validate:"required"`
	UserID         uint64 `json:"user_id" validate:"required"`
	PostCategoryID uint64 `json:"post_category_id" validate:"required"`
}

type PostUpdateRequest struct {
	Title          string `json:"title" validate:"required,max=255"`
	Code           string `json:"code" validate:"required,max=100"`
	Description    string `json:"description" validate:"required"`
	PostCategoryID uint64 `json:"post_category_id" validate:"required"`
}

type PostGetRequest struct {
	Code string `json:"code" validate:"required"`
}

type PostDeleteRequest struct {
	Code string `json:"code" validate:"required"`
}
