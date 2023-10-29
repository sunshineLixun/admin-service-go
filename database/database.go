package database

import (
	"admin-service-go/config"
	"admin-service-go/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDb() (err error) {

	if err := godotenv.Load(".env.local"); err != nil {
		log.Panic("读取配置失败, 错误原因", err.Error())
	}

	var (
		user     = config.Config("DB_USER")
		password = config.Config("DB_PASSWORD")
		host     = config.Config("DB_HOST")
		db       = config.Config("DB_NAME")
		_port    = config.Config("DB_PORT")
	)

	port, err := strconv.Atoi(_port)

	if err != nil {
		return err
	}

	// dsn := "root:123456789@tcp(127.0.0.1:3306)/admin_service_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db)

	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("链接失败, 错误原因", err)
		return err
	}

	sqlDB, err := DBConn.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("链接数据库成功")

	DBConn.AutoMigrate(&models.User{})

	return nil

}
