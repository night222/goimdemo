package initialiaze

import (
	"context"
	"fmt"
	"goimdemo/common"
	"time"

	"github.com/go-redis/redis/v8"
)

// 初始化redis连接池
func InitRedis() (err error) {
	address := fmt.Sprintf("%s:%s", common.ConfigData.RedisCofig.Host, common.ConfigData.RedisCofig.Port)
	common.RedisClient = redis.NewClient(&redis.Options{
		Addr:         address,
		Network:      "tcp",
		Password:     common.ConfigData.RedisCofig.Password,
		DialTimeout:  time.Duration(common.ConfigData.ConnectTimeOut) * time.Second,
		ReadTimeout:  time.Duration(common.ConfigData.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(common.ConfigData.WriteTimeOut) * time.Second,
		IdleTimeout:  time.Duration(common.ConfigData.RedisCofig.IdleTimeout) * time.Second,
		PoolSize:     common.ConfigData.RedisCofig.PoolSize,
		MinIdleConns: common.ConfigData.RedisCofig.MinIdleConns,
	})
	_, err = common.RedisClient.Ping(context.TODO()).Result()
	return
}
