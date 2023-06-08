package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"-"`
	Username      string         `gorm:"column:username;type:varchar(64);not null;uniqueIndex" json:"username"` // 用户名
	Password      string         `gorm:"column:password;type:varchar(128);not null" json:"-"`                   // 密码
	Email         string         `gorm:"column:email;type:varchar(256);not null;unique" json:"email"`           // 邮件地址
	EmailVerified bool           `gorm:"column:email_verified;not null;default:false" json:"email_verified"`
	Role          string         `gorm:"column:role;type:varchar(32)" json:"role"`
	LoginIP       string         `gorm:"column:login_ip;type:varchar(64);not null;default:''" json:"login_ip"`  // 最近登录IP
	LoginUA       string         `gorm:"column:login_ua;type:varchar(256);not null;default:''" json:"login_ua"` // 最近登录UA
	LoginAt       time.Time      `gorm:"column:login_at;default:current_timestamp" json:"login_at"`             // 最近登录时间
	CreatedAt     time.Time      `gorm:"column:created_at;default:current_timestamp" json:"-"`                  // 创建时间
	DeletedAt     gorm.DeletedAt `json:"-"`                                                                     // 软删除
}

type LoginInfo struct {
	Username  string `binding:"required"`
	Password  string `binding:"required"`
	IP        string
	UserAgent string
}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
