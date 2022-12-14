package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`        // 用户名
	Avatar   string `gorm:"type:varchar(255);not null"`       // 头像
	UserNick string `gorm:"type:varchar(50);not null"`        // 昵称
	Phone    string `gorm:"type:varchar(11);not null;unique"` // 手机号
	Password string `gorm:"size:255;not null"`                // 密码
	LoginIp  string `gorm:"type:varchar(20);not null"`        // 登录IP
	Email    string `gorm:"size:255;not null"`                // 邮箱
	Group    string `gorm:"size:255;not null"`                // 角色组 1,2,3,4
	LoginAt  string `gorm:"size:255;not null"`                // 登录时间
}
