package model

import "github.com/golang-jwt/jwt/v5"

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserResponse
}

// type Error struct {
// 	Field    string   `json:"field"`
// 	Messages []string `json:"messages"`
// }
