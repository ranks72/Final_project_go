package service

import (
	"final_project_go/entity"
	"final_project_go/pkg/helpers"
	"final_project_go/repository/photo_repository"
	"final_project_go/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	photoRepo photo_repository.PhotoRepository
}

func NewAuthService(
	userRepo user_repository.UserRepository,
	photoRepo photo_repository.PhotoRepository,
) AuthService {
	return &authService{
		userRepo:  userRepo,
		photoRepo: photoRepo,
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
				"err_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		photoIdParam, err := helpers.GetParamId(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "not found",
			})
			return
		}

		if photo.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}
