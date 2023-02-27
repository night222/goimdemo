package utils

import (
	"goimdemo/common"
)

// @Title:close.go
// @Author:proket
// @Description: 关闭连接
// @Since:2023-02-25 00:01:49
func Close() {
	common.RedisClient.Close()
}
