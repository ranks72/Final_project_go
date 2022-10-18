package photo_repository

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
)

type PhotoRepository interface {
	CreatePhotoRepo(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotoRepo() ([]entity.Photo, errs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
	EditedPhoto(photoId int, photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletedPhoto(photoId int) error
}
