package comment_pg

import (
	"errors"
	"final_project_go/entity"
	"final_project_go/pkg/errs"
	"final_project_go/repository/comment_repository"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (u *commentPG) GetCommentById(commentId int) (*entity.Comment, errs.MessageErr) {
	comment := &entity.Comment{}
	err := u.db.First(comment, "id", commentId).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return comment, nil
}

func (u *commentPG) CreateCommentRepo(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	if err := u.db.Create(commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return commentPayload, nil
}

func (u *commentPG) GetAllCommentRepo() ([]entity.Comment, errs.MessageErr) {
	payloadComment := []entity.Comment{}
	if err := u.db.Preload("User").Preload("Photo").Find(&payloadComment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}
	return payloadComment, nil
}

func (u *commentPG) EditedCommentRepo(commentId int, commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	query := u.db.Where("id", commentId).Updates(commentPayload)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return nil, errs.NewNotFoundError("comments doesn't exit")
	}

	return commentPayload, nil
}

func (u *commentPG) DeletedCommentRepo(commentId int) error {
	query := u.db.Delete(new(*entity.Comment), "id", commentId)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return errors.New("NOT FOUND")
	}

	return err
}
