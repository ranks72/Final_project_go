package user_pg

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/repository/user_repository"

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
		First(user, "email", userPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return user, nil
}

func (u *userPG) Register(userPayload *entity.User) errs.MessageErr {
	err := u.db.Create(userPayload).Error

	if err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	return nil
}

func (u *userPG) GetUserByIdAndEmail(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	var user entity.User
	// row := u.db.QueryRow(retrieveUserByIdAndEmail, userPayload.Id, userPayload.Email)

	// err := row.Scan(&user.Id, &user.Email)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, errs.NewNotFoundError("users doesn't exit")
	// 	}
	// 	return nil, errs.NewInternalServerErrorr("something went wrong")
	// }

	return &user, nil
}
