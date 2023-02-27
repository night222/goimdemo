package common

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// @Title:redis.go
// @Author:proket
// @Description: redis的公共方法和对象
// @Since:2023-02-24 23:26:17
var RedisClient *redis.Client

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
