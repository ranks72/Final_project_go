package sosmed_pg

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/repository/sosmed_repository"

	"gorm.io/gorm"
)

type sosmedPG struct {
	db *gorm.DB
}

func NewSosmedPG(db *gorm.DB) sosmed_repository.SosmedRepository {
	return &sosmedPG{
		db: db,
	}
}

func (u *sosmedPG) CreateSosmedRepo(sosmedPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	if err := u.db.Create(sosmedPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return sosmedPayload, nil
}

func (u *sosmedPG) GetSosmedById(sosmedId int) (*entity.SocialMedia, errs.MessageErr) {
	sosmed := &entity.SocialMedia{}
	err := u.db.First(sosmed, "id", sosmedId).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return sosmed, nil
}

func (u *sosmedPG) GetAllSosmedRepo() ([]entity.SocialMedia, errs.MessageErr) {
	payloadSosmed := []entity.SocialMedia{}
	if err := u.db.Preload("User").Find(&payloadSosmed).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return payloadSosmed, nil
}
