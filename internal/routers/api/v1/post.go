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

	posts, err := services.GetPosts(queryPost)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}

	appG.Response(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	post, err := services.GetPost(postId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, post)
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
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, post)
}

func ArchivePost(c *gin.Context) {
	var (
		appG   = app.Gin{C: c}
		postId = appG.C.Param("postId")
	)

	post, err := services.ArchivePost(postId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, post)
}

func UpdatePostStatus(c *gin.Context) {
	var (
		appG             = app.Gin{C: c}
		updatePostStatus = models.UpdatePostStatus{}
	)

	if err := appG.C.ShouldBindJSON(&updatePostStatus); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	post, err := services.UpdatePostStatus(updatePostStatus)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, post)
}
