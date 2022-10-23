package dto

import (
	"final_project_go/entity"
	"time"
)

type CommentRequest struct {
	Message string `json:"message" valid:"required~Message of your comment is required" example:"example"`
	PhotoID int    `json:"photo_id" valid:"required~PhotoID of your comment is required" example:"1"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Message of your comment is required" example:"example"`
}

type CommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message" binding:"required" example:"example"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateCommentResponses(data *entity.Comment) CommentResponse {
	return CommentResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

type GetCommentResponse struct {
	ID        int                  `json:"id"`
	Message   string               `json:"message" binding:"required" example:"example"`
	PhotoID   int                  `json:"photo_id"`
	UserID    int                  `json:"user_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	User      UserResponseComment  `json:"user"`
	Photo     PhotoCommentResponse `json:"photo"`
}

func ObjectAllcomment(data entity.Comment) GetCommentResponse {
	return GetCommentResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: UserResponseComment{
			ID:       data.User.ID,
			Email:    data.User.Email,
			Username: data.User.Username,
		},
		Photo: PhotoCommentResponse{
			ID:       data.Photo.ID,
			Title:    data.Photo.Title,
			Caption:  data.Photo.Caption,
			PhotoUrl: data.Photo.PhotoUrl,
			UserID:   data.Photo.UserID,
		},
	}
}

func GetAllCommentResponse(res []entity.Comment) (responses []GetCommentResponse) {
	for _, comment := range res {
		responses = append(responses, ObjectAllcomment(comment))
	}
	return
}

func UpdateCommentResponses(data *entity.Comment) CommentResponse {
	return CommentResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		UpdatedAt: data.UpdatedAt,
	}
}
