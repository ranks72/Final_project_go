package photo_pg

import (
	"errors"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/repository/photo_repository"

	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (u *photoPG) CreatePhotoRepo(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	if err := u.db.Create(photoPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return photoPayload, nil
}

func (u *photoPG) GetAllPhotoRepo() ([]entity.Photo, errs.MessageErr) {
	payloadPhoto := []entity.Photo{}
	if err := u.db.Preload("User").Find(&payloadPhoto).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return payloadPhoto, nil
}

func (u *photoPG) EditedPhoto(photoId int, photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	query := u.db.Where("id", photoId).Updates(photoPayload)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return nil, errs.NewNotFoundError("photo doesn't exit")
	}

	return photoPayload, nil
}

func (u *photoPG) GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr) {
	photo := &entity.Photo{}
	err := u.db.First(photo, "id", photoId).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return photo, nil
}

func (u *photoPG) DeletedPhoto(photoId int) error {
	query := u.db.Delete(new(*entity.Photo), "id", photoId)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return errors.New("NOT FOUND")
	}

	return err
}
