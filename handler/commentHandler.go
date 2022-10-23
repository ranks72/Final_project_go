package handler

import (
	"final_project_go/dto"
	"final_project_go/entity"
	"final_project_go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentRestHandler struct {
	service service.CommentService
}

func newCommentHandler(commentService service.CommentService) commentRestHandler {
	return commentRestHandler{
		service: commentService,
	}
}

func (u commentRestHandler) AddCommentHandler(c *gin.Context) {
	var commentRequest dto.CommentRequest

	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.PostComment(userData.ID, &commentRequest)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, dto.CreateCommentResponses(result))
}

func (u commentRestHandler) GetAllCommentHandler(c *gin.Context) {
	// userData := c.MustGet("userData")

	// fmt.Printf("userData => %+v", userData)
	result, err := u.service.GetAllComment()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//c.JSON(http.StatusOK, result)
	c.JSON(http.StatusOK, dto.GetAllCommentResponse(result))
}

func (u commentRestHandler) UpdatedCommentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	var commentRequest dto.UpdateCommentRequest
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	result, err := u.service.UpdatedComment(id, &commentRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"msg": err.Error(),
			"err": "Comment not found",
		})
		return
	}
	c.JSON(http.StatusOK, dto.UpdateCommentResponses(result))

}

func (u commentRestHandler) DeletedCommentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	delete := u.service.DeletedComment(id)
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

	c.JSON(http.StatusOK, "your comment has been successfully deleted")

}
