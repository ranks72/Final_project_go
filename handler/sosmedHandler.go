package handler

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sosmedRestHandler struct {
	service service.SosmedService
}

func newSosmedHandler(sosmedService service.SosmedService) sosmedRestHandler {
	return sosmedRestHandler{
		service: sosmedService,
	}
}

func (u sosmedRestHandler) AddSosmedHandler(c *gin.Context) {
	var sosmedRequest dto.SosmedRequest

	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	if err := c.ShouldBindJSON(&sosmedRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.PostSosmed(userData.ID, &sosmedRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Message(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	_ = result

	c.JSON(http.StatusCreated, dto.CreateSosmedResponses(result))
}
