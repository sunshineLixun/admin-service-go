package models

import "gorm.io/gorm"

// InputRole swagger输入参数
type InputRole struct {
	RoleName string `gorm:"type:varchar(20);not null;unique" json:"roleName" validate:"required"`
}

type Role struct {
	gorm.Model

	InputRole
	// 关联用户
	Users []User `gorm:"many2many:user_roles;joinForeignKey:roleId;joinReferences:userId;" json:"users"`
}

type ResponseRole struct {
	ID       uint   `json:"roleId"`
	RoleName string `json:"roleName"`
}

type UpdateRoleInput struct {
	RoleName string `json:"roleName" validate:"required"`
}
