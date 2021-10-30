package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type GetJinsMemeTokenRoot struct {
	Value []GetJinsMemeTokenRes `json:"value"`
}

type SaveJinsMemeDataBind struct {
	UserID     uint               `json:"user_id"`
	ConcDataID primitive.ObjectID `json:"conc_id" bson:"_conc_id"`
	StartTime  time.Time          `json:"start_time" bson:"start_time"`
	EndTime    time.Time          `json:"end_time" bson:"end_time"`
}

type SaveJinsMemeDataReq struct {
	AccessToken string `json:"access_token"`
	StartTime   string `json:"start_time" bson:"start_time"`
	EndTime     string `json:"end_time" bson:"end_time"`
}

type SaveJinsMemeDataRes struct {
	ComputedData interface{} `json:"computed_data"`
	Cursor       string      `json:"cursor"`
}

type SaveJinsMemeDataSave struct {
	SaveJinsMemeDataRes SaveJinsMemeDataRes `json:"jins_meme_data"`
	ConcDataID          primitive.ObjectID  `json:"conc_id" bson:"_conc_id"`
}
