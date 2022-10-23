package handler

import (
	"final_project_go/database"
	"final_project_go/repository/comment_repository/comment_pg"
	"final_project_go/repository/photo_repository/photo_pg"
	"final_project_go/repository/sosmed_repository/sosmed_pg"
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

	//comment
	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo, userRepo, photoRepo)
	commentRestHandler := newCommentHandler(commentService)

	//sosmed
	sosmedRepo := sosmed_pg.NewSosmedPG(db)
	sosmedService := service.NewSosmedService(sosmedRepo)
	sosmedRestHandler := newSosmedHandler(sosmedService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, sosmedRepo)

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
	photoRoute := router.Group("/photos")
	{
		photoRoute.Use(authService.Authentication())
		photoRoute.GET("", photoRestHandler.GetAllPhotoHandler)
		photoRoute.POST("", photoRestHandler.AddPhotoHandler)
		photoRoute.PUT("/:photoId", authService.PhotoAuthorization(), photoRestHandler.UpdatedPhotoHandler)
		photoRoute.DELETE("/:photoId", authService.PhotoAuthorization(), photoRestHandler.DeletedPhotoHandler)
	}

	commentRoute := router.Group("/comments")
	{
		commentRoute.Use(authService.Authentication())
		commentRoute.POST("", commentRestHandler.AddCommentHandler)
		commentRoute.GET("", commentRestHandler.GetAllCommentHandler)
		commentRoute.PUT("/:commentId", authService.CommentAuthorization(), commentRestHandler.UpdatedCommentHandler)
		commentRoute.DELETE("/:commentId", authService.CommentAuthorization(), commentRestHandler.DeletedCommentHandler)
	}

	sosmedRoute := router.Group("/socialmedias")
	{
		sosmedRoute.Use(authService.Authentication())
		sosmedRoute.POST("", sosmedRestHandler.AddSosmedHandler)
		sosmedRoute.GET("", sosmedRestHandler.GetAllSosmedHandler)
		sosmedRoute.PUT("/:sosmedId", authService.SocialMediaAuthorization(), sosmedRestHandler.UpdatedSosmedHandler)
		sosmedRoute.DELETE("/:sosmedId", authService.SocialMediaAuthorization(), sosmedRestHandler.DeletedSosmedHandler)
	}

	fmt.Println("Server running on PORT =>", PORT)
	router.Run(PORT)
	return router
}
