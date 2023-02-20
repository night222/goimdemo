package config

import "gopkg.in/ini.v1"

type AppConfig struct {
	DatabaseConfig `ini:"database"`
	MD5Salt        `ini:"md5_salt"`
}

//数据库配置
type DatabaseConfig struct {
	DatabaseName string `ini:"database_name"`
	Host         string `ini:"host"`
	Port         string `ini:"port"`
	UserName     string `ini:"user_name"`
	Passworld    string `ini:"passworld"`
	Charset      string `ini:"charset"`
	ParseTime    string `ini:"parse_time"`
	Supplement   string `ini:"supplement"`
}

//设置md5加密
type MD5Salt struct {
	Max uint `ini:"max"`
}

func InitConfig() (apc AppConfig, err error) {
	err = ini.MapTo(&apc, "./config/app.ini")
	return
}
