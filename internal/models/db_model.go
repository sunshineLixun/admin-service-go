package models

import (
	"admin-service-go/global"
	"admin-service-go/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDbEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	s := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s, databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("链接失败, 错误原因", err)
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	log.Println("链接数据库成功")

	return db, nil

}

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = NewDbEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}
