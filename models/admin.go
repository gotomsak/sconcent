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
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// UserSignin userのサインイン時のstruct
type AdminUserSignin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSend Signup時に送られるdata
type AdminUserSend struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// AdminGetIDLogsRes AdminGetIdLogsで返されるデータ
type AdminGetIDLogsRes struct {
	GetIDLogs []GetIDLog `json:"get_id_logs"`
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
