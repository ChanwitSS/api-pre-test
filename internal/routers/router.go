package routers

import (
	// _ "post/docs"

	v1 "post/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/post", v1.GetPosts)
		apiv1.GET("/post/:postId", v1.GetPost)
		apiv1.POST("/post", v1.CreatePost)
		apiv1.POST("/post/archive/:postId", v1.ArchivePost)
		// apiv1.POST("/post/update-status", v1.UpdatePostStatus)
		// apiv1.PATCH("/post/:postId", v1.UpdatePost)
		// apiv1.DELETE("/post/:postId", v1.DeletePost)

		apiv1.POST("/comment", v1.CreateComment)
	}

	return r
}
