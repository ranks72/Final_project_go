package handler

import (
	"final_project_go/dto"
	"final_project_go/service"
	"net/http"
	"strconv"
	"strings"

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

	result, err := u.service.Register(&user)

	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"msg": err.Error(),
					"err": "email telah digunakan",
				})
				return
			}

			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"msg": err.Error(),
				"err": "username telah digunakan",
			})
			return
		}
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, dto.DataRegisterResponse(*result))
}

func (u userRestHandler) Updated(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var user dto.UpdateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.UpdatedUser(id, &user)
	if err != nil {
		if err.Error() == "users doesn't exit" {
			c.JSON(http.StatusNotFound, map[string]interface{}{
				"msg": http.StatusText(http.StatusNotFound),
			})
			return
		}

		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"msg": err.Error(),
					"err": "email telah digunakan",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": err.Error(),
				"err": "username telah digunakan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "BAD_REQUEST",
		})
		return
	}
	c.JSON(http.StatusOK, dto.DataUpdateResponse(*result))
}

func (u userRestHandler) Deleted(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	delete := u.service.DeletedUser(id)
	if delete != nil {
		if delete.Error() == "NOT FOUND" {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": err.Error(),
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
