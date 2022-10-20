package handler

import (
	"final_project_go/database"
	"final_project_go/repository/photo_repository/photo_pg"
	"final_project_go/repository/user_repository/user_pg"
	"final_project_go/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func StartServer() *gin.Engine {

	router := gin.Default()
	db := database.GetDb()

	//user
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := newUserHandler(userService)

	//photo
	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := newPhotoHandler(photoService)

	authService := service.NewAuthService(userRepo, photoRepo)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	userRoute := router.Group("/users")
	{
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.PUT("", authService.Authentication(), userRestHandler.Updated)
		userRoute.DELETE("", authService.Authentication(), userRestHandler.Deleted)
	}
	photoRoute := router.Group("/photo")
	{
		photoRoute.Use(authService.Authentication())
		photoRoute.GET("", photoRestHandler.GetAllPhotoHandler)
		photoRoute.POST("", photoRestHandler.AddPhotoHandler)
		photoRoute.PUT("/:photoId", authService.PhotoAuthorization(), photoRestHandler.UpdatedPhotoHandler)
		photoRoute.DELETE("/:photoId", authService.PhotoAuthorization(), photoRestHandler.DeletedPhotoHandler)
	}

	fmt.Println("Server running on PORT =>", PORT)
	router.Run(PORT)
	return router
}
