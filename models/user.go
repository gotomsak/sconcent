package models

import (
	"github.com/jinzhu/gorm"
)

// User userテーブルのstruct
type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"size:255"`
	Email          string `json:"email" gorm:"type:varchar(100);unique_index"`
	PasswordDigest string `json:"password_digest" gorm:"size:255"`
}

// UserSignup userのサインアップ時のstruct
type UserSignup struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// UserSignin userのサインイン時のstruct
type UserSignin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSend Signup時に送られるdata
type UserSend struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// UserInfo passwordを除いたuser情報
type UserInfo struct {
	// gorm.Model
	UserID   uint   `json:"user_id"`
	Username string `json:"username" gorm:"size:255"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
}
