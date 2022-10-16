package user_repository

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
)

type UserRepository interface {
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	Login(user *entity.User) (*entity.User, errs.MessageErr)
	Register(user *entity.User) errs.MessageErr
	GetUserByIdAndEmail(user *entity.User) (*entity.User, errs.MessageErr)
}
