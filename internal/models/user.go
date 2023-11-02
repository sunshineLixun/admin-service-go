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
