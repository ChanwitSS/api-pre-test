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
		appG    = app.Gin{C: c}
		orderId = appG.C.Param("orderId")
	)

	order, err := services.GetPost(orderId)
	if err != nil {
		appG.C.JSON(http.StatusInternalServerError, err)
		return
	}
	appG.C.JSON(http.StatusOK, order)
}
