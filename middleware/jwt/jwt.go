package jwt

// import (
// 	"rooster-blog/pkg/e"
// 	"rooster-blog/pkg/setting"

// 	"net/http"
// 	"time"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// var jwtSecret = []byte(setting.JwtSecret)

// type Claims struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	jwt.StandardClaims
// }

// func GenerateToken(username,password string) (string,error) {
// 	nowTime := time.Now()
// 	expireTime := nowTime.Add(3 * time.Hour)

// 	claims := Claims{
// 		username,
// 		password,
// 		jwt.StandardClaims{
// 			ExpiresAt: expireTime.Unix(),
// 			Issuer: "rooster-blog",
// 		},
// 	}

// 	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
// 	token,e := tokenClaims.SignedString(jwtSecret)
// 	return token,e
// }

// func ParseToken(token string) (*Claims,error) {
// 	tokenClaims,e := jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret,nil
// 	})
// 	if tokenClaims != nil {
// 		if claim,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid {
// 			return claim,nil
// 		}
// 	}
// 	return nil,e
// }

// func JWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var code int
// 		var data interface{}

// 		code = e.SUCCESS
// 		token := c.Query("token")
// 		if token == "" {
// 			code = e.INVALID_PARAMS
// 		} else {
// 			claims,err := ParseToken(token)
// 			if err != nil {
// 				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
// 			} else if time.Now().Unix() > claims.ExpiresAt {
// 				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
// 			}
// 		}

// 		if code != e.SUCCESS {
// 			c.JSON(http.StatusUnauthorized,gin.H{
// 				"code":code,
// 				"msg":e.GetMsg(code),
// 				"data":data,
// 			})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }