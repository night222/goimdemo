package config

import "gopkg.in/ini.v1"

type AppConfig struct {
	DatabaseConfig `ini:"database"`
	MD5Salt        `ini:"md5_salt"`
	Log            `ini:"log"`
	Token          `ini:"token"`
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

//记录日志配置
type Log struct {
	Path         string `ini:"path"`
	RotationTime int64  `ini:"rotation_time"`  //日志切割时间 单位 minute
	RotationSize int64  `init:"rotation_size"` //日志文件大小 单位 kb
}

//设置md5加密
type MD5Salt struct {
	Max uint `ini:"max"`
}

//token
type Token struct {
	Singin    string `ini:"singin"`
	ExpiredAt uint   `ini:"expired_at"` // minute
	Subject   string `ini:"subject"`
}

func InitConfig() (apc AppConfig, err error) {
	err = ini.MapTo(&apc, "./config/app.ini")
	return
}
