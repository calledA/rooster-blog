package router

import (
	"github.com/gin-gonic/gin"
	"rooster-blog/middleware/jwt"
	"rooster-blog/api/article"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	articleApi := router.Group("/api/admin")
	articleApi.Use(jwt.JWT())
	{
		articleApi.GET("/articles",article.GetArticles)
	}
	return router
}
