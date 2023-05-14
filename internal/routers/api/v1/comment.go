package v1

import (
	"net/http"
	"post/internal/models"
	"post/internal/services"
	"post/pkg/app"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var (
		appG          = app.Gin{C: c}
		createComment = models.Comment{}
	)

	if err := appG.C.ShouldBindJSON(&createComment); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	comment, err := services.CreateComment(createComment)
	if err != nil {
		appG.C.JSON(http.StatusInternalServerError, app.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: err,
		})
		return
	}
	appG.C.JSON(http.StatusOK, app.Response{
		Code: http.StatusOK,
		Data: comment,
	})
}
