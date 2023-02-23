package config

import "gopkg.in/ini.v1"

type AppConfig struct {
	DatabaseConfig `ini:"database"`
	MD5Salt        `ini:"md5_salt"`
	Log            `ini:"log"`
	Token          `ini:"token"`
	RedisCofig     `ini:"redis"`
	Group          `ini:"group"`
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
	Level        string `ini:"level"`          //记录日志级别最小
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

type RedisCofig struct {
	Host           string `ini:"host"`
	Port           string `ini:"port"`
	ConnectTimeOut uint   `ini:"connect_time_out"` // 秒
	ReadTimeOut    uint   `ini:"read_time_out"`    // 秒
	WriteTimeOut   uint   `ini:"write_time_out"`   // 秒
	MinIdleConns   int    `ini:"min_idle_conns"`   //最大空闲连接数
	PoolSize       int    `ini:"pool_size"`        // 表示和数据库的最大连接数，0表示没有限制
	IdleTimeout    uint   `ini:"idle_time_out"`    //最大空闲时间 秒
	Password       string `ini:"password"`
}

//更具id分组
type Group struct {
	UserId       []string `ini:"user_id"`
	GroupId      []string `ini:"group_id"`
	PostfixUser  string   `ini:"postfix_user"`
	PostfixGroup string   `ini:"postfix_group"`
}

func InitConfig() (apc AppConfig, err error) {
	err = ini.MapTo(&apc, "./config/app.ini")
	return
}
