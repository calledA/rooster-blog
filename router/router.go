package router

import (
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
		adminApi.GET("/click",articles.GetArticleClicksApi)

		//文章点击数排名
		adminApi.GET("/rank",articles.GetArticleRankApi)

		//总的主题排名
		adminApi.GET("/topic_rank",articles.GetTopicRankApi)
	}

	

	return router
}
