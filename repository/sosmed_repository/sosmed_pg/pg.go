package sosmed_pg

import (
	"errors"
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

func (u *sosmedPG) EditedSosmedRepo(sosmedId int, sosmedPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	query := u.db.Where("id", sosmedId).Updates(sosmedPayload)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return nil, errs.NewNotFoundError("social media doesn't exit")
	}

	return sosmedPayload, nil
}

func (u *sosmedPG) DeletedSosmedRepo(sosmedId int) error {
	query := u.db.Delete(new(*entity.SocialMedia), "id", sosmedId)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return errors.New("NOT FOUND")
	}

	return err
}
