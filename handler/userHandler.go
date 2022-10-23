package handler

import (
	"final_project_go/dto"
	"final_project_go/entity"
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
			"msg": err.Message(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (u userRestHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest
	//var errormsg helpers.ErrorMsg

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": err.Error(),
		})
		return
	}

	result, err := u.service.Register(&user)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, dto.DataRegisterResponse(*result))
}

func (u userRestHandler) Updated(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	var user dto.UpdateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.UpdatedUser(userData.ID, &user)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, dto.DataUpdateResponse(*result))
}

func (u userRestHandler) Deleted(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	delete := u.service.DeletedUser(userData.ID)
	if delete != nil {
		if delete.Error() == "NOT FOUND" {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": delete.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "BAD_REQUEST",
		})
		return
	}
	c.JSON(http.StatusOK, "your account has been successfully deleted")
}
