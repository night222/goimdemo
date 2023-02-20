package utils

import (
	"math/rand"
	"time"
)

// 生成随机字符串
func RandString(max uint) string {
	rand.Seed(time.Now().Unix())
	len := rand.Intn(int(max))
	dataSlice := make([]byte, len)
	for i := 0; i < len; i++ {
		temNum := rand.Intn(95)
		dataSlice[i] = '!' + byte(temNum)
	}
	return string(dataSlice)
}
