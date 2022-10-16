package handler

import (
	"final_project_go/database"
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

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := newUserHandler(userService)

	authService := service.NewAuthService(userRepo)
	_ = authService

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	userRoute := router.Group("/users")
	{
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.PUT("/:userId")
		userRoute.DELETE("/:userId")
	}

	fmt.Println("Server running on PORT =>", PORT)
	router.Run(PORT)
	return router
}
