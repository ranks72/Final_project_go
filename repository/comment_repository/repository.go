package comment_repository

import (
	"final_project_go/entity"
	"final_project_go/pkg/errs"
)

type CommentRepository interface {
	GetCommentById(commentId int) (*entity.Comment, errs.MessageErr)
	CreateCommentRepo(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetAllCommentRepo() ([]entity.Comment, errs.MessageErr)
	EditedCommentRepo(commentId int, commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	DeletedCommentRepo(commentId int) error
}
