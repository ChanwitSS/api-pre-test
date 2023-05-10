package v1

import (
	"net/http"
	"post/internal/models"
	"post/internal/services"
	"post/pkg/app"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var (
		appG      = app.Gin{C: c}
		queryPost = models.QueryPost{}
		// queryOrder = models.QueryPost{}
	)

	if err := c.ShouldBindQuery(&queryPost); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	orders, err := services.GetPosts(queryPost)
	if err != nil {
		appG.C.JSON(http.StatusInternalServerError, err)
		return
	}

	appG.C.JSON(http.StatusOK, orders)
}

func GetPost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	order, err := services.GetPost(postId)
	if err != nil {
		appG.C.JSON(http.StatusInternalServerError, err)
		return
	}
	appG.C.JSON(http.StatusOK, order)
}

func ArchivePost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	post, err := services.ArchivePost(postId)
	if err != nil {
		appG.C.JSON(http.StatusInternalServerError, err)
		return
	}
	appG.C.JSON(http.StatusOK, post)
}
