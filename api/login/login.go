package login

import (
	"net/http"
	"rooster-blog/middleware/jwt"
	"rooster-blog/models"
	"rooster-blog/pkg/e"
	"rooster-blog/pkg/logging"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Email    string
	Password string
}

func Login(ctx *gin.Context) {
	var admin Admin
	var code int
	err := ctx.BindJSON(&admin)
	if err != nil {
		logging.Fatal("登陆绑定JSON失败", err)
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&admin)

	if ok {
		isExist := models.CheckAdmin(admin.Email, admin.Password)
		if isExist {
			token, err := jwt.GenerateToken(admin.Email, admin.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				ctx.SetCookie("Authorization", token, 7*24*60*60, "/", "127.0.0.1", false, true)
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	if code == 200 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   e.GetMsg(code),
		})
	}

}
