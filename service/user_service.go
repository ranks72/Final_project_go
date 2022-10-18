package service

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/user_repository"
	"fmt"
)

type UserService interface {
	Register(userPayload *dto.RegisterRequest) (*entity.User, errs.MessageErr)
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdatedUser(userId int, user *dto.UpdateRequest) (*entity.User, errs.MessageErr)
	DeletedUser(userId int) error
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*entity.User, errs.MessageErr) {
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

	data_user, err := u.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	return data_user, nil
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

func (u *userService) UpdatedUser(userId int, userPayload *dto.UpdateRequest) (*entity.User, errs.MessageErr) {

	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	user, err := u.userRepo.EditedUser(userId, payload)

	if err != nil {
		return nil, err
	}

	user, err = u.userRepo.GetUserById(userId)
	fmt.Println(user)

	return user, nil
}

func (u *userService) DeletedUser(userId int) error {
	user := u.userRepo.DeletedUser(userId)

	if user != nil {
		return nil
	}
	return user
}
