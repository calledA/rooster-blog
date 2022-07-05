package router

import (
	"rooster-blog/api/article"
	"rooster-blog/api/login"
	"rooster-blog/middleware/cors"
	"rooster-blog/middleware/jwt"
	"rooster-blog/models"
	_ "rooster-blog/pkg/logging"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Cors())
	
	router.POST("/login",login.Login)
	router.GET("/home",models.Oauth)

	articleApi := router.Group("/api/admin")
	articleApi.Use(jwt.JWT())
	{
		articleApi.GET("/articles",article.GetArticles)
	}
	return router
}
