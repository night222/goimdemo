// @Title:system_init.go
// @Author:proket
// @Description: 初始化文件
// @Since:2023-02-18 17:24:17
package utils

import (
	"context"
	"fmt"
	"goimdemo/config"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ConfigData config.AppConfig
var DB *gorm.DB
var RedisClient *redis.Client

// 初始化
func Init() (err error) {
	//初始化config
	err = InitConfig()
	if err != nil {
		return
	}
	fmt.Println("config app init success")
	//初始化log
	err = InitLoger("/self_log_%Y-%m-%d.log")
	if err != nil {
		return
	}
	fmt.Println("log init successs")
	//初始化mysql连接
	err = InitMysql()
	if err != nil {
		return
	}
	fmt.Println("Mysql init success")
	//初始化redis连接池
	err = InitRedis()
	if err != nil {
		return
	}
	fmt.Println("Redis init success")
	InitMessage()
	fmt.Println("message init success")
	ginLogWriter, err := RotationFile("/gin_log_%Y-%m-%d.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = io.MultiWriter(ginLogWriter, os.Stdout)
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
		SlowThreshold:             time.Second, //慢sql阈值
		LogLevel:                  logger.Info, //级别
		Colorful:                  true,        //彩色
		IgnoreRecordNotFoundError: false,       //是否显示没有查询到的结果
	})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s%s",
		ConfigData.UserName,
		ConfigData.Passworld,
		ConfigData.DatabaseConfig.Host,
		ConfigData.DatabaseConfig.Port,
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

// 初始化redis连接池
func InitRedis() (err error) {
	address := fmt.Sprintf("%s:%s", ConfigData.RedisCofig.Host, ConfigData.RedisCofig.Port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         address,
		Network:      "tcp",
		Password:     ConfigData.RedisCofig.Password,
		DialTimeout:  time.Duration(ConfigData.ConnectTimeOut) * time.Second,
		ReadTimeout:  time.Duration(ConfigData.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(ConfigData.WriteTimeOut) * time.Second,
		IdleTimeout:  time.Duration(ConfigData.RedisCofig.IdleTimeout) * time.Second,
		PoolSize:     ConfigData.RedisCofig.PoolSize,
		MinIdleConns: ConfigData.RedisCofig.MinIdleConns,
	})
	_, err = RedisClient.Ping(context.TODO()).Result()
	return
}
func InitMessage() {
	GroupByUserId = ConfigData.Group.UserId
	GroupByUserId = ConfigData.Group.GroupId
	MessageQueueChan = make(map[string]map[uint]chan Message, len(GroupByUserId))
	for i := 0; i < len(GroupByUserId); i++ {
		MessageQueueChan[GroupByUserId[i]] = make(map[uint]chan Message, 1000)
	}
}

// 关闭连接
func Close() {
	RedisClient.Close()
}
