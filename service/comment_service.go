package service

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/comment_repository"
	"final_project_go/repository/photo_repository"
	"final_project_go/repository/user_repository"
	"time"
)

type CommentService interface {
	PostComment(userId int, commentPayload *dto.CommentRequest) (*entity.Comment, errs.MessageErr)
	GetAllComment() ([]entity.Comment, errs.MessageErr)
	UpdatedComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*entity.Comment, errs.MessageErr)
	DeletedComment(commentId int) error
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
	userRepo    user_repository.UserRepository
	photoRepo   photo_repository.PhotoRepository
}

func NewCommentService(
	commentRepo comment_repository.CommentRepository,
	userRepo user_repository.UserRepository,
	photoRepo photo_repository.PhotoRepository,
) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		photoRepo:   photoRepo,
	}
}

func (u *commentService) PostComment(userId int, commentPayload *dto.CommentRequest) (*entity.Comment, errs.MessageErr) {

	err := helpers.ValidateStruct(commentPayload)
	if err != nil {
		return nil, err
	}

	payloadPhoto := &entity.Comment{
		UserID:    userId,
		Message:   commentPayload.Message,
		PhotoID:   commentPayload.PhotoID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	photo, err := u.photoRepo.GetPhotoById(commentPayload.PhotoID)
	_ = photo

	if err != nil {
		return nil, err
	}

	comment, err := u.commentRepo.CreateCommentRepo(payloadPhoto)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (u *commentService) GetAllComment() ([]entity.Comment, errs.MessageErr) {
	comment, err := u.commentRepo.GetAllCommentRepo()
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (u *commentService) UpdatedComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*entity.Comment, errs.MessageErr) {
	err := helpers.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.Comment{
		Message:   commentPayload.Message,
		UpdatedAt: time.Now(),
	}
	comment, err := u.commentRepo.EditedCommentRepo(commentId, payload)
	if err != nil {
		return nil, err
	}

	comment, err = u.commentRepo.GetCommentById(commentId)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (u *commentService) DeletedComment(commentId int) error {
	comment := u.commentRepo.DeletedCommentRepo(commentId)

	if comment != nil {
		return nil
	}
	return comment
}
