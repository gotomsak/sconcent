package models

import "github.com/jinzhu/gorm"

type GetJinsMemeTokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type GetJinsMemeTokenReq struct {
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
	RedirectUri  string `json:"redirect_uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GetJinsMemeTokenBind struct {
	UserID uint   `json:"user_id"`
	Code   string `json:"code"`
}

type GetJinsMemeTokenSave struct {
	gorm.Model
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
}
