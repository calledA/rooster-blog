package router

import (
	"github.com/gin-gonic/gin"
	"rooster-blog/models"
	"rooster-blog/pkg/logging"
	"rooster-blog/middleware/jwt"
	"rooster-blog/api/article"
	"rooster-blog/pkg/cors"

)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Cors())
	
	logging.Info("开始请求")
	
	router.GET("/login",models.CheckAuth)
	router.GET("/home",models.Oauth)

	articleApi := router.Group("/api/admin")
	articleApi.Use(jwt.JWT())
	{
		articleApi.GET("/articles",article.GetArticles)
	}
	return router
}
