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
	Users []User `gorm:"many2many:user_roles" json:"users"`
}

// UserRoles
// 当使用多个字段作为主键时，GORM 默认会使用这些字段的组合作为主键。这意味着在UserRole关联表中，(UserID, RoleID)将作为复合主键。
// 如果您只想使用单个字段作为主键，可以使用gorm:"primaryKey;autoIncrement:false"标签来禁用自增功能。
type UserRoles struct {
	UserID uint `gorm:"primaryKey" json:"user_id"`
	RoleID uint `gorm:"primaryKey" json:"role_id"`
}
