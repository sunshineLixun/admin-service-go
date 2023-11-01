package global

import (
	"admin-service-go/pkg/setting"
	"time"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

func SetupSetting() error {
	set, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = set.ReadSection("Server", ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", DatabaseSetting)
	if err != nil {
		return err
	}

	ServerSetting.ReadTimeout *= time.Second

	ServerSetting.WriteTimeout *= time.Second

	return nil

}
