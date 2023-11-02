package models

type ResponseHTTP struct {
	Success bool        `json:"success" gorm:"default:true"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code" gorm:"default:1"`
}
