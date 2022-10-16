package dto

import "final_project_go/entity"

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
	Age      int    `json:"age" valid:"required~age cannot be empty"`
}

type RegisterResponse struct {
	//Message  string `json:"message"`
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func MapToRegisterResponse(data entity.User) RegisterResponse {
	return RegisterResponse{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
		Age:      data.Age,
	}
}
