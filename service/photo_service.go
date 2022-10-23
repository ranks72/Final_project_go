package service

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/photo_repository"
)

type PhotoService interface {
	PostPhoto(userId int, photoPayload *dto.RequestPhoto) (*entity.Photo, errs.MessageErr)
	GetAllPhoto() ([]entity.Photo, errs.MessageErr)
	FindPhotoid(photoID int) (*entity.Photo, errs.MessageErr)
	UpdatedPhoto(photoId int, photoPayload *dto.RequestPhoto) (*entity.Photo, errs.MessageErr)
	DeletedPhoto(photoId int) error
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (u *photoService) PostPhoto(userId int, photoPayload *dto.RequestPhoto) (*entity.Photo, errs.MessageErr) {

	err := helpers.ValidateStruct(photoPayload)
	if err != nil {
		return nil, err
	}

	payloadPhoto := &entity.Photo{
		UserID:   userId,
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
	}

	photo, err := u.photoRepo.CreatePhotoRepo(payloadPhoto)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *photoService) GetAllPhoto() ([]entity.Photo, errs.MessageErr) {
	photos, err := u.photoRepo.GetAllPhotoRepo()
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (u *photoService) FindPhotoid(photoId int) (*entity.Photo, errs.MessageErr) {
	photo, err := u.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *photoService) UpdatedPhoto(photoId int, photoPayload *dto.RequestPhoto) (*entity.Photo, errs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
	}
	photo, err := u.photoRepo.EditedPhoto(photoId, payload)
	if err != nil {
		return nil, err
	}
	photo, err = u.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *photoService) DeletedPhoto(photoId int) error {
	photo := u.photoRepo.DeletedPhoto(photoId)

	if photo != nil {
		return nil
	}
	return photo
}
