package router

import (
	// "rooster-blog/api/article"
	"rooster-blog/api/login"
	"rooster-blog/api/admin/articles"
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
		//总点击数
		adminApi.GET("/clicks",articles.GetClicks)

		//文章点击数排名
		adminApi.GET("/ranks",login.LoginAuth)

		//总访问量(按月排)
		adminApi.GET("/visits",login.LoginAuth)

		//每天的访问主题数量
		adminApi.GET("/topic_visit",login.LoginAuth)

		//总的主题排名
		adminApi.GET("/topic_rank",login.LoginAuth)
	}

	return router
}
