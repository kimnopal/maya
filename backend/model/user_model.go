package model

type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
