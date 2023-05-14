package routers

import (
	// _ "post/docs"

	"post/internal/routers/api"
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

	auth := r.Group("/auth")
	{
		auth.POST("register", api.Register)
		auth.POST("login", api.Login)
	}

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{
		post := apiv1.Group("/post")
		{
			post.GET("", v1.GetPosts)
			post.GET("/:postId", v1.GetPost)
			post.POST("", v1.CreatePost)
			post.POST("/archive/:postId", v1.ArchivePost)
			// post.POST("/post/update-status", v1.UpdatePostStatus)
			// post.PATCH("/post/:postId", v1.UpdatePost)
			// post.DELETE("/post/:postId", v1.DeletePost)
		}

		comment := apiv1.Group("/comment")
		{
			comment.POST("", v1.CreateComment)
		}

	}

	return r
}
