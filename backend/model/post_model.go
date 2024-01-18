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
