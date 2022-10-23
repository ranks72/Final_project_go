package dto

import (
	"final_project_go/entity"
	"time"
)

type SosmedRequest struct {
	Name           string `json:"name" valid:"required~Name of your Social Media is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Url of Social Media is required"`
}

// respon
type CreateSosmedResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func CreateSosmedResponses(data *entity.SocialMedia) CreateSosmedResponse {
	return CreateSosmedResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		CreatedAt:      data.CreatedAt,
	}
}
