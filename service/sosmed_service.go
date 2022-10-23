package service

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/sosmed_repository"
	"time"
)

type SosmedService interface {
	PostSosmed(userId int, sosmedPayload *dto.SosmedRequest) (*entity.SocialMedia, errs.MessageErr)
	GetAllSosmed() ([]entity.SocialMedia, errs.MessageErr)
	UpdatedSosmed(sosmedId int, sosmedPayload *dto.SosmedRequest) (*entity.SocialMedia, errs.MessageErr)
	DeletedSosmed(sosmedId int) error
}

type sosmedService struct {
	sosmedRepo sosmed_repository.SosmedRepository
}

func NewSosmedService(sosmedRepo sosmed_repository.SosmedRepository) SosmedService {
	return &sosmedService{
		sosmedRepo: sosmedRepo,
	}
}

func (u *sosmedService) PostSosmed(userId int, sosmedPayload *dto.SosmedRequest) (*entity.SocialMedia, errs.MessageErr) {
	err := helpers.ValidateStruct(sosmedPayload)
	if err != nil {
		return nil, err
	}

	payloadSosmed := &entity.SocialMedia{
		UserID:         userId,
		Name:           sosmedPayload.Name,
		SocialMediaUrl: sosmedPayload.SocialMediaUrl,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	sosmed, err := u.sosmedRepo.CreateSosmedRepo(payloadSosmed)

	if err != nil {
		return nil, err
	}

	return sosmed, nil
}

func (u *sosmedService) GetAllSosmed() ([]entity.SocialMedia, errs.MessageErr) {
	sosmeds, err := u.sosmedRepo.GetAllSosmedRepo()
	if err != nil {
		return nil, err
	}

	return sosmeds, nil
}

func (u *sosmedService) UpdatedSosmed(sosmedId int, sosmedPayload *dto.SosmedRequest) (*entity.SocialMedia, errs.MessageErr) {
	err := helpers.ValidateStruct(sosmedPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.SocialMedia{
		Name:           sosmedPayload.Name,
		SocialMediaUrl: sosmedPayload.SocialMediaUrl,
	}

	sosmed, err := u.sosmedRepo.EditedSosmedRepo(sosmedId, payload)
	if err != nil {
		return nil, err
	}

	sosmed, err = u.sosmedRepo.GetSosmedById(sosmedId)

	if err != nil {
		return nil, err
	}

	return sosmed, nil
}

func (u *sosmedService) DeletedSosmed(sosmedId int) error {
	sosmed := u.sosmedRepo.DeletedSosmedRepo(sosmedId)

	if sosmed != nil {
		return nil
	}
	return sosmed
}
