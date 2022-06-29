package login

type UserInfo struct {
	AccessToken  string `json:"access_token"`
	ExpireIn     string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

type WxUserInfo struct {
	NickName string
	
}
