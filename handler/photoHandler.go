package handler

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoRestHandler struct {
	service service.PhotoService
}

func newPhotoHandler(photoService service.PhotoService) photoRestHandler {
	return photoRestHandler{
		service: photoService,
	}
}

// @Tags Photo
// @Summary Post photo
// @ID post-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.RequestPhoto true "json request body"
// @Success 201 {object} dto.CreatePhotoResponse
// @Router /photos [post]
func (u photoRestHandler) AddPhotoHandler(c *gin.Context) {
	var photoRequest dto.RequestPhoto

	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.PostPhoto(userData.ID, &photoRequest)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, dto.CreatePhotoResponses(result))
}

// @Tags Photo
// @Summary Get all photos
// @ID get-all-photos
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} dto.PhotoResponse
// @Router /photos [get]
func (u photoRestHandler) GetAllPhotoHandler(c *gin.Context) {
	// userData := c.MustGet("userData")

	// fmt.Printf("userData => %+v", userData)
	result, err := u.service.GetAllPhoto()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//c.JSON(http.StatusOK, result)
	c.JSON(http.StatusOK, dto.GetAllPhotoResponse(result))
}

// @Tags Photo
// @Summary Update photo
// @ID update-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "photoId"
// @Param RequestBody body dto.RequestPhoto true "json request body"
// @Success 200 {object} dto.UpdateResponse
// @Router /photos/{photoId} [put]
func (u photoRestHandler) UpdatedPhotoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	var photoRequest dto.RequestPhoto
	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}
	result, err := u.service.UpdatedPhoto(id, &photoRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}
	c.JSON(http.StatusOK, dto.UpdatedPhotoResponse(*result))
}

// @Tags Photo
// @Summary Delete photo
// @ID delete-photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "photoId"
// @Success 200 {object} dto.DeletePhotoResponse
// @Router /photos/{photoId} [delete]
func (u photoRestHandler) DeletedPhotoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	delete := u.service.DeletedPhoto(id)
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

	res := dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	c.JSON(http.StatusOK, res)
}
