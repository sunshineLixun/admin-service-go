package models

import "gorm.io/gorm"

type UserSwagger struct {
	UserName string `gorm:"not null" json:"userName" validate:"required"`
	Password string `gorm:"not null" json:"password" validate:"required"`
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

// UpdateUserInput 更新用户输入参数
type UpdateUserInput struct {
	UserName string `json:"userName" validate:"required"`
}
