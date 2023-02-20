// @Title:system_init.go
// @Author:proket
// @Description: 初始化文件
// @Since:2023-02-18 17:24:17
package utils

import (
	"fmt"
	"goimdemo/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ConfigData config.AppConfig
var DB *gorm.DB

// 初始化
func Init() (err error) {
	//初始化config
	err = InitConfig()
	if err != nil {
		return
	}
	fmt.Println("config app init")
	//初始化mysql连接
	err = InitMysql()
	if err != nil {
		return
	}
	fmt.Println("Mysql init")
	return
}

// 初始化配置文件
func InitConfig() (err error) {
	ConfigData, err = config.InitConfig()
	return
}

// 初始化mysql连接
func InitMysql() (err error) {
	//设置把sql日志输出到控制台
	newLOger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢sql阈值
		LogLevel:      logger.Info, //级别
		Colorful:      true,        //彩色
	})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s%s",
		ConfigData.UserName,
		ConfigData.Passworld,
		ConfigData.Host,
		ConfigData.Port,
		ConfigData.DatabaseName,
		ConfigData.Charset,
		ConfigData.ParseTime,
		ConfigData.Supplement,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLOger,
	})
	return
}
