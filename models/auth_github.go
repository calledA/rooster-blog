package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rooster-blog/pkg/logging"

	"github.com/gin-gonic/gin"
)


type Auth struct {
	ClientId string
	ClientSecret string
	RedirectUrl string
}

var auth = Auth{
	ClientId :"747198252bd3eb94eebe",
	ClientSecret:"2f0342e8afbde6692d52611f4df52d63e94c2df5",
	RedirectUrl:"http://localhost:9100/home",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}


func CheckAuth(ctx *gin.Context){
	ctx.Redirect(http.StatusMovedPermanently,fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&client_id=%v",auth.ClientId,auth.RedirectUrl))
}

func Oauth(ctx *gin.Context) {
	var err error
	var token *Token
	code := ctx.Query("code")
	tokenUrl := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",auth.ClientId,auth.ClientSecret,code)
	if token,err = GetToken(tokenUrl);err != nil {
		logging.Error("获取token失败",err)
		return
	}

	var userInfo map[string]interface{}

	if userInfo,err = GetUserInfo(token);err != nil {
		logging.Error("获取用户信息失败:",err)
		return
	}

	for k, v := range userInfo {
		logging.Error("userinfo",k,v)
	}

	fmt.Printf("%+v",token)
}

// 获取 token
func GetToken(url string) (*Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		logging.Error("形成请求失败",err)
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		logging.Error("获得响应失败",err)
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		logging.Error("解析Token失败",err)
		return nil, err
	}
	return &token, nil
}

// 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {

	// 形成请求
	var userInfoUrl = "https://api.github.com/user"	// github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		logging.Error("获取用户信息失败",err)
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	
	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		logging.Error("获取用户信息响应失败",err)
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logging.Error("数据写入 userInfo 失败",err)
		return nil, err
	}
	return userInfo, nil
}


