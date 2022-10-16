package handler

import (
	"final_project_go/dto"
	"final_project_go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	service service.UserService
}

func newUserHandler(userService service.UserService) userRestHandler {
	return userRestHandler{
		service: userService,
	}
}

func (u userRestHandler) Login(c *gin.Context) {
	var user dto.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	token, err := u.service.Login(&user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusCreated, token)
}

func (u userRestHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	successMessage, err := u.service.Register(&user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}
