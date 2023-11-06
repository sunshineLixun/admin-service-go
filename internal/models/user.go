package models

import "gorm.io/gorm"

type UserSwagger struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model

	UserSwagger
}

// ResponseUser 用于返回用户信息，去掉密码字段
type ResponseUser struct {
	gorm.Model
	UserName string `json:"userName"`
}
