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

type GetSosmedResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserResponsePhoto
}

func ObjectAllsosmeds(data entity.SocialMedia) GetSosmedResponse {
	return GetSosmedResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		User: UserResponsePhoto{
			Email:    data.User.Email,
			Username: data.User.Username,
		},
	}
}

func GetAllSosmedResponse(res []entity.SocialMedia) (responses []GetSosmedResponse) {
	for _, sosmed := range res {
		responses = append(responses, ObjectAllsosmeds(sosmed))
	}
	return
}

type UpdateSosmedResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func UpdatedsosmedsResponse(data entity.SocialMedia) UpdateSosmedResponse {
	return UpdateSosmedResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		UpdatedAt:      data.UpdatedAt,
	}
}
