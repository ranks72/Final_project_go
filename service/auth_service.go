package service

import (
	"final_project_go/entity"
	"final_project_go/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
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
