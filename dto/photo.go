package dto

import (
	"final_project_go/entity"
	"time"
)

// request
type RequestPhoto struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

// func (data *RequestPhoto) RequestPhotos(userId int) *entity.Photo {
// 	return &entity.Photo{
// 		UserID:   userId,
// 		Title:    data.Title,
// 		Caption:  data.Caption,
// 		PhotoUrl: data.PhotoUrl,
// 	}
// }

// respons
type CreatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
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
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	PhotoUrl  string            `json:"photo_url"`
	UserID    int               `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
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
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
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
