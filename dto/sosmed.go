package dto

import (
	"final_project_go/entity"
	"time"
)

type SosmedRequest struct {
	Name           string `json:"name" valid:"required~Name of your Social Media is required" example:"example"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Url of Social Media is required" example:"https://instagram.com"`
}

// respon
type CreateSosmedResponse struct {
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"example"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://instagram.com"`
	UserID         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
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
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"example"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://instagram.com"`
	UserID         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
	UpdatedAt      time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
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
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"example"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://instagram.com"`
	UserID         int       `json:"user_id" example:"1"`
	UpdatedAt      time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
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

type DeleteSosmedResponse struct {
	Message string `json:"message" example:"your social media has been successfully deleted"`
}
