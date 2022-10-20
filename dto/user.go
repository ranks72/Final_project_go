package dto

import (
	"final_project_go/entity"
	"time"
)

//request

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type RegisterRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty" binding:"required"`
	Email    string `json:"email" valid:"required~email cannot be empty,email~Invalid email format" binding:"required,email"`
	Password string `json:"password" valid:"required~password cannot be empty" binding:"required,min=6"`
	Age      int    `json:"age" valid:"required~age cannot be empty" binding:"required,min=9"`
}

type UpdateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

// response
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func DataRegisterResponse(data entity.User) RegisterResponse {
	return RegisterResponse{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
		Age:      data.Age,
	}
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DataUpdateResponse(data entity.User) UpdateResponse {
	return UpdateResponse{
		ID:        data.ID,
		Email:     data.Email,
		Username:  data.Username,
		Age:       data.Age,
		UpdatedAt: data.UpdatedAt,
	}
}

type UserResponsePhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
