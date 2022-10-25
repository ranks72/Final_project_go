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

// @Tags Social Media
// @Summary Create social media
// @ID create-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.SosmedRequest true "json request body"
// @Success 201 {object} dto.CreateSosmedResponse
// @Router /socialmedias [post]
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

// @Tags Social Media
// @Summary Get all social medias
// @ID get-all-social-medias
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} dto.GetSosmedResponse
// @Router /socialmedias [get]
func (u sosmedRestHandler) GetAllSosmedHandler(c *gin.Context) {
	result, err := u.service.GetAllSosmed()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//c.JSON(http.StatusOK, result)
	c.JSON(http.StatusOK, dto.GetAllSosmedResponse(result))
}

// @Tags Social Media
// @Summary Update social media
// @ID update-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "sosmedId"
// @Param RequestBody body dto.SosmedRequest true "json request body"
// @Success 200 {object} dto.UpdateSosmedResponse
// @Router /socialmedias/{sosmedId} [put]
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

// @Tags Social Media
// @Summary Delete social media
// @ID delete-social-media
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "sosmedId"
// @Success 200 {object} dto.DeleteSosmedResponse
// @Router /socialmedias/{sosmedId} [delete]
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

	res := dto.DeleteSosmedResponse{
		Message: "Your sosmed has been successfully deleted",
	}
	c.JSON(http.StatusOK, res)
}
