package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `json:"userNmae"`

	Password string `json:"password"`
}
