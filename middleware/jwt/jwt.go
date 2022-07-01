package jwt

import (
	"fmt"
	// "go/token"
	"rooster-blog/models"
	"rooster-blog/pkg/e"
	"rooster-blog/pkg/logging"

	// "rooster-blog/pkg/setting"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("23347$040412")


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Admin struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("JWT方法")
		var code int
		var data interface{}
		var admin Admin

		code = e.SUCCESS
		tokenString := c.Request.Header.Get("Authorization")
		username := c.Query("username")
		password := c.Query("password")

		if tokenString == "" {
			models.DB.Select("id").Where(Admin{Username: username,Password: password}).First(&admin)
			if admin.ID < 0 {
				return
			}
			token,err := GenerateToken(username,password)
			if err != nil {
				logging.Error("生成token出错",err)
				return
			}
			c.SetCookie("Authorization", token, 7*24*60*60, "/", "127.0.0.1", false, true)
		} else {
			claims,err := ParseToken(tokenString)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GenerateToken(username,password string) (string,error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "rooster-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,e := tokenClaims.SignedString(jwtSecret)
	return token,e
}

func ParseToken(token string) (*Claims,error) {
	tokenClaims,e := jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token) (interface{}, error) {
		return jwtSecret,nil
	})
	if tokenClaims != nil {
		if claim,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid {
			return claim,nil
		}
	}
	return nil,e
}


// func RefreshToken(tokenString string) *Claims {
// 	jwt.TimeFunc = func() time.Time {
// 		return time.Unix(0,0)
// 	}
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	claims,ok := token.Claims.(*Claims)
// 	if !ok {
// 		logging.Fatal("token生成失败")
// 	}
// 	jwt.TimeFunc = time.Now
// 	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
// 	return GenerateToken(claims)
// }
