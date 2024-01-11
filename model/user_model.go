package model

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,max:100"`
	Name     string `json:"name" validate:"required,max:100"`
	Email    string `json:"email" validate:"required,max:100"`
	Password string `json:"password" validate:"required,max:100"`
	MajorID  uint64 `json:"major_id" validate:"required"`
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	MajorID  uint64 `json:"major_id"`
}
