package dto

import (
	"final_project_go/entity"
	"time"
)

//request

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty" example:"Testing@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty" example:"Passowrd"`
}

type RegisterRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty" binding:"required" example:"Testing"`
	Email    string `json:"email" valid:"required~email cannot be empty,email~Invalid email format" binding:"required,email" example:"Testing@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty, password~min=9" binding:"required,min=6" example:"Passowrd"`
	Age      int    `json:"age" valid:"required~age cannot be empty,age~min=9" binding:"required,min=9" example:"10"`
}

type UpdateRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty,email~Invalid email format" binding:"required,email" example:"Testing@gmail.com"`
	Username string `json:"username" valid:"required~username cannot be empty" binding:"required" example:"TestingLagi"`
}

// response
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type RegisterResponse struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"Testing"`
	Email    string `json:"email" example:"Testing@gmail.com"`
	Age      int    `json:"age" example:"10"`
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
	ID        int       `json:"id" example:"1"`
	Email     string    `json:"email" example:"Testing@gmail.com"`
	Username  string    `json:"username" example:"TestingLagi"`
	Age       int       `json:"age" example:"10"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
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

type DeleteResponse struct {
	Message string `json:"message" example:"your account has been successfully deleted"`
}

type UserResponsePhoto struct {
	Email    string `json:"email" example:"Testing@gmail.com"`
	Username string `json:"username" example:"TestingLagi"`
}

type UserResponseComment struct {
	ID       int    `json:"id" example:"1"`
	Email    string `json:"email" example:"Testing@gmail.com"`
	Username string `json:"username" example:"TestingLagi"`
}
