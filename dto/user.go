package dto

import (
	"time"
)

type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required~Full name cannot be ampty"`
	Email    string `json:"email" valid:"email,required~Email cannot be ampty"`
	Password string `json:"password" valid:"required~Password cannot be ampty,minstringlength(6)~password min 6 length"`
}

type RegisterResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"email,required~Email cannot be ampty"`
	Password string `json:"password" valid:"required~Password cannot be ampty,minstringlength(6)~password min 6 length" `
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"required~Full name cannot be ampty"`
	Email    string `json:"email" valid:"email,required~Email cannot be ampty"`
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
