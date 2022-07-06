package router

import (
	// "rooster-blog/api/article"
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
	
	router.GET("/home",models.Oauth)
	router.POST("/api/admin/login",login.LoginAuth)

	adminApi := router.Group("/api/admin")
	adminApi.Use(jwt.JWT())
	{
		//管理员登陆
		adminApi.POST("/articles",login.LoginAuth)

	}

	return router
}
