package handler

import (
	"errors"
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type photoRestHandler struct {
	service service.PhotoService
}

func newPhotoHandler(photoService service.PhotoService) photoRestHandler {
	return photoRestHandler{
		service: photoService,
	}
}

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
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			errormsg := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				errormsg[i] = ErrorMsg{getErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, errormsg)
			return
		}

		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.PostPhoto(userData.ID, &photoRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Message(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreatePhotoResponses(result))
}

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

	c.JSON(http.StatusOK, "your photo has been successfully deleted")
}
