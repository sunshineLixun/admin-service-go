package models

import "gorm.io/gorm"

type UserLogin struct {
	UserName string `gorm:"not null" json:"userName" validate:"required"`
	Password string `gorm:"not null" json:"password" validate:"required"`
}

type CreateUserInput struct {
	UserLogin
	RoleIds []uint `gorm:"-" json:"roleIds"`
}

type UpdateUserInput struct {
	UserName string `json:"userName" validate:"required"`
}

type RoleModel struct {
	Roles []Role `gorm:"many2many:user_roles;joinForeignKey:userId;joinReferences:roleId" json:"roles"`
}

type User struct {
	gorm.Model

	CreateUserInput

	RoleModel
}

// ResponseUser 用于返回用户信息，去掉密码字段
type ResponseUser struct {
	ID       uint           `json:"userId"`
	UserName string         `json:"userName"`
	Roles    []ResponseRole `json:"roles"`
}
