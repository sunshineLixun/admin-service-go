package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDb() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/admin_service_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("链接失败, 错误原因", err)
		os.Exit(2)
	}

	log.Println("链接成功")

	db.AutoMigrate()

	DBConn = db
}
