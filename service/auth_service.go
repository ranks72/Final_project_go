package service

import (
	"final_project_go/entity"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/comment_repository"
	"final_project_go/repository/photo_repository"
	"final_project_go/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
	//SocialMediaAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.UserRepository
	photoRepo   photo_repository.PhotoRepository
	commentRepo comment_repository.CommentRepository
}

func NewAuthService(
	userRepo user_repository.UserRepository,
	photoRepo photo_repository.PhotoRepository,
	commentRepo comment_repository.CommentRepository,
) AuthService {
	return &authService{
		userRepo:    userRepo,
		photoRepo:   photoRepo,
		commentRepo: commentRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var user *entity.User = &entity.User{}

		tokenStr := ctx.Request.Header.Get("Authorization")

		err := user.VerifyToken(tokenStr)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": err.Error(),
			})
			return
		}

		_, err = a.userRepo.GetUserByIdAndEmail(user)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": err.Error(),
			})
			return
		}

		ctx.Set("userData", *user)
		ctx.Next()
	})
}

func (a *authService) PhotoAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Bad Request",
				"message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		photoIdParam, err := helpers.GetParamId(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if photo.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	})
}

func (a *authService) CommentAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		_ = userData

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Bad Request",
				"message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		commentIdParam, err := helpers.GetParamId(ctx, "commentId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		comment, err := a.commentRepo.GetCommentById(commentIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if comment.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		photo, err := a.photoRepo.GetPhotoById(comment.PhotoID)
		_ = photo

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "your comment in photo doesn't exist",
			})
			return
		}

		ctx.Next()
	})
}

// func (a *authService) SocialMediaAuthorization() gin.HandlerFunc {
// 	return gin.HandlerFunc(func(ctx *gin.Context) {

// 	})
// }
