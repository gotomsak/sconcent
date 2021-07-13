package models

import (
	"github.com/jinzhu/gorm"
)

// User userテーブルのstruct
type AdminUser struct {
	gorm.Model
	Username       string `json:"username" gorm:"size:255"`
	Email          string `json:"email" gorm:"type:varchar(100);unique_index"`
	PasswordDigest string `json:"password_digest" gorm:"size:255"`
}

// UserSignup userのサインアップ時のstruct
type AdminUserSignup struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// UserSignin userのサインイン時のstruct
type AdminUserSignin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// UserSend Signup時に送られるdata
type AdminUserSend struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// AdminGetIDLogsRes AdminGetIdLogsで返されるデータ
type AdminGetIDLogAllRes struct {
	GetIDLogAll []GetIDLog `json:"get_id_log_all"`
}

type AdminGetIDLogUserRes struct {
	GetIDLogUser []GetIDLog `json:"get_id_log_user"`
}

// GetUserAllRes
type AdminGetUserAllRes struct {
	// gorm.Model
	UsersInfo []UserInfo `json:"users_info"`
}

// // AdminGetRecAllRes
// type AdminGetRecAllRes struct {

// 	ConcentData []GetConcentrationRes `json:"concent_data"`
// }
