package articles

import (
	"log"
	"rooster-blog/models"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetClicks(ctx *gin.Context) {
	// idStr := ctx.DefaultQuery("id","1")
	idStr := ctx.Query("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}
	topic := models.GetClicks(id)

	ctx.JSON(http.StatusOK, gin.H{
		// "code":  code,
		// "msg":   e.GetMsg(code),
		"data":topic,
	})
}