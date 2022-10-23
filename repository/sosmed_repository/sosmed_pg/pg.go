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
