package articles

import (
	_ "fmt"
	"rooster-blog/models"
	"rooster-blog/pkg/e"
	"rooster-blog/pkg/logging"

	// "strconv"
	// "fmt"
	"net/http"

	// "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type Rank struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Click int    `json:"click"`
}

//总点击数
func GetArticleClicksApi(ctx *gin.Context) {

	// idStr := ctx.Query("id")
	// id,err := strconv.Atoi(idStr)
	var code int
	var count int

	article, err := models.GetClicks()
	if err != nil {
		code = e.ERROR_QUERY
		logging.Fatal(err)
	} else {
		code = e.SUCCESS
	}

	for _, val := range article {
		count += val.ArticleClick
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"total": count,
	})
}

//文章点击数排名
func GetArticleRankApi(ctx *gin.Context) {
	var code int
	rank := make([]interface{}, 0)

	article, err := models.GetArticleRank()
	if err != nil {
		code = e.ERROR_QUERY
		logging.Fatal(err)
		// return nil, err
	} else {
		code = e.SUCCESS
	}

	for _, val := range article {
		rank = append(rank, Rank{Id: val.Id, Title: val.Title, Click: int(val.ArticleClick)})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": rank,
	})

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

//总的主题排名
func GetTopicRankApi(ctx *gin.Context) {
	var code int
	rank := make([]interface{}, 0)

	topic, err := models.GetTopicRank()
	if err != nil {
		code = e.ERROR_QUERY
		logging.Fatal(err)
	} else {
		code = e.SUCCESS
	}

	for _, val := range topic {
		rank = append(rank, Rank{Id: val.Id, Title: val.Name, Click: int(val.TagClick)})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": rank,
	})
}
