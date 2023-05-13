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
		appG.Response(http.StatusBadRequest, app.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: strings.Split(err.Error(), "\n"),
		})
		return
	}

	posts, err := services.GetPosts(queryPost)
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
		Data: posts,
	})
}

func GetPost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	post, err := services.GetPost(postId)
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
		Data: post,
	})
}

func CreatePost(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		createPost = models.Post{}
	)

	if err := appG.C.ShouldBindJSON(&createPost); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	post, err := services.CreatePost(createPost)
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
		Data: post,
	})
}

func ArchivePost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	post, err := services.ArchivePost(postId)
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
		Data: post,
	})
}
