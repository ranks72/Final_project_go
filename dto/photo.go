package dto

import (
	"final_project_go/entity"
	"time"
)

// request
type RequestPhoto struct {
	Title    string `json:"title" valid:"required~Title of your photo is required" example:"Photo Posting"`
	Caption  string `json:"caption" example:"Coba coba"`
	PhotoUrl string `json:"photo_url" valid:"required~Url of your photo cannot be empty" example:"www.google.img.com"`
}

// respons
type CreatePhotoResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Photo Posting"`
	Caption   string    `json:"caption" example:"Coba coba"`
	PhotoUrl  string    `json:"photo_url" example:"www.google.img.com"`
	UserID    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func CreatePhotoResponses(data *entity.Photo) CreatePhotoResponse {
	return CreatePhotoResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

type PhotoResponse struct {
	ID        int               `json:"id" example:"1"`
	Title     string            `json:"title" example:"Photo Posting"`
	Caption   string            `json:"caption" example:"Coba coba"`
	PhotoUrl  string            `json:"photo_url" example:"www.google.img.com"`
	UserID    int               `json:"user_id" example:"1"`
	CreatedAt time.Time         `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
	UpdatedAt time.Time         `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
	User      UserResponsePhoto `json:"user"`
}

func ObjectAllphotos(data entity.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: UserResponsePhoto{
			Email:    data.User.Email,
			Username: data.User.Username,
		},
	}
}

func GetAllPhotoResponse(res []entity.Photo) (responses []PhotoResponse) {
	for _, photo := range res {
		responses = append(responses, ObjectAllphotos(photo))
	}
	return
}

type UpdatePhoto struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Photo Posting"`
	Caption   string    `json:"caption" example:"Coba coba"`
	PhotoUrl  string    `json:"photo_url" example:"www.google.img.com"`
	UserID    int       `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func UpdatedPhotoResponse(data entity.Photo) UpdatePhoto {
	return UpdatePhoto{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		UpdatedAt: data.UpdatedAt,
	}
}

type DeletePhotoResponse struct {
	Message string `json:"message" example:"your photo has been successfully deleted"`
}

type PhotoCommentResponse struct {
	ID       int    `json:"id"  example:"1"`
	Title    string `json:"title" example:"Photo Posting"`
	Caption  string `json:"caption" example:"Coba coba"`
	PhotoUrl string `json:"photo_url" example:"www.google.img.com"`
	UserID   int    `json:"user_id" example:"1"`
}
