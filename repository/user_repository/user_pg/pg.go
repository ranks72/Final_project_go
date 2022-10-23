package user_pg

import (
	"errors"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/repository/user_repository"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {

	user := &entity.User{}

	err := u.db.First(user, "id", userId).Error

	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil

}

func (u *userPG) Login(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := &entity.User{}

	if err := u.db.Select("id", "email", "password").
		First(user, "email", userPayload.Email).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("email dan username salah")
	}
	return user, nil
}

func (u *userPG) Register(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	fmt.Println(userPayload)
	if err := u.db.Create(userPayload).Error; err != nil {
		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				checkerr := errs.NewInternalServerErrorr("email is already used")
				return nil, checkerr
			}
			checkerr := errs.NewInternalServerErrorr("username is already used")
			return nil, checkerr
		}
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return userPayload, nil
}

func (u *userPG) GetUserByIdAndEmail(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := &entity.User{}

	if err := u.db.Select("id", "email").
		First(user, "id", userPayload.ID).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}

func (u *userPG) EditedUser(id int, userPayload *entity.User) (*entity.User, errs.MessageErr) {
	query := u.db.Where("id", id).Updates(userPayload)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				checkerr := errs.NewInternalServerErrorr("email is already used")
				return nil, checkerr
			}
			checkerr := errs.NewInternalServerErrorr("username is already used")
			return nil, checkerr
		}
	}

	user := &entity.User{}

	err = u.db.First(user, "id", id).Error

	if err != nil {
		return nil, errs.NewInternalServerErrorr("akun tidak ditemukan")
	}

	return userPayload, nil
}

func (u *userPG) DeletedUser(id int) error {
	query := u.db.Delete(new(*entity.User), "id", id)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return errors.New("NOT FOUND")
	}

	return err
}
