package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpHost     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         int
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
}

func (s *Setting) ReadSection(key string, value interface{}) error {
	if err := s.vp.UnmarshalKey(key, value); err != nil {
		return err
	}
	return nil
}
