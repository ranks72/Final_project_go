package handler

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/service"
	"net/http"
	"strconv"

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
		c.JSON(err.Status(), err)
		return
	}

	_ = result

	c.JSON(http.StatusCreated, dto.CreateSosmedResponses(result))
}

func (u sosmedRestHandler) GetAllSosmedHandler(c *gin.Context) {
	result, err := u.service.GetAllSosmed()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//c.JSON(http.StatusOK, result)
	c.JSON(http.StatusOK, dto.GetAllSosmedResponse(result))
}

func (u sosmedRestHandler) UpdatedSosmedHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sosmedId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	var sosmedRequest dto.SosmedRequest
	if err := c.ShouldBindJSON(&sosmedRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}
	result, err := u.service.UpdatedSosmed(id, &sosmedRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "Name and Social Media URL Tidak boleh kosong",
		})
		return
	}
	c.JSON(http.StatusOK, dto.UpdatedsosmedsResponse(*result))
}

func (u sosmedRestHandler) DeletedSosmedHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sosmedId"))
	if err != nil && id < 1 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	delete := u.service.DeletedSosmed(id)
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

	c.JSON(http.StatusOK, "your sosmed has been successfully deleted")
}
