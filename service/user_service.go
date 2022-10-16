package service

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/user_repository"
)

type UserService interface {
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Username: userPayload.Username,
		Age:      userPayload.Age,
	}

	err = user.HashPass()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		//Message:  "user data has been successfully created",
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Username: userPayload.Username,
		Age:      userPayload.Age,
	}

	return res, nil
}

func (u *userService) Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Email: userPayload.Email,
	}

	user, err := u.userRepo.Login(payload)

	if err != nil {

		return nil, err
	}

	validPassword := user.ComparePassword(userPayload.Password)

	if !validPassword {
		return nil, errs.NewNotAuthenticated("invalid email/password")
	}

	token := user.GenerateToken()

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}
