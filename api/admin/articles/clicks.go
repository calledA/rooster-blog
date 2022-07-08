package articles

import (
	"rooster-blog/models"
	"rooster-blog/pkg/e"
	"rooster-blog/pkg/logging"

	// "strconv"
	// "fmt"
	"net/http"

	// "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

//总点击数
func GetArticleClicksApi(ctx *gin.Context) {
	
	// idStr := ctx.Query("id")
	// id,err := strconv.Atoi(idStr)
	var code int
	count, err := models.GetClicks()
	if err != nil {
		code = e.ERROR_QUERY
		logging.Fatal(err)
	} else {
		code = e.SUCCESS
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": count,
	})
}

//文章点击数排名
func GetArticleRankApi(ctx *gin.Context) {
	models.GetArticleRank()


	//redis
	// conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	logging.Fatal("redis connect err", err)
	// }
	// defer conn.Close()
	// //执行redis
	// _, err = conn.Do("ZADD", "mykey", "INCR", 1, "robot2")
	// if err != nil {
	// 	logging.Fatal("redis插入错误", err)
	// 	return
	// }
	// tests, err := redis.StringMap(conn.Do("ZRANGE", "mykey", 0, 10, "withscores"))
	// if err != nil {
	// 	logging.Fatal("redis err", err)
	// 	return
	// }
	// for test := range tests {
	// 	fmt.Println(tests[test])
	// }

}

//总访问量(按月排)
func GetArticleVisitApi(ctx *gin.Context) {

}

//每天的访问主题数量
func GetTopicVisitApi(ctx *gin.Context) {

}

//总的主题排名
func GetTopicRankApi(ctx *gin.Context) {

}
