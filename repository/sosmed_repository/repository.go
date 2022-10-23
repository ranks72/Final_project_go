package sosmed_repository

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
)

type SosmedRepository interface {
	CreateSosmedRepo(sosmedPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
}
