package initialiaze

import (
	"fmt"
	"goimdemo/utils"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

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
	ginLogWriter, err := utils.RotationFile("/gin_log_%Y-%m-%d.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = io.MultiWriter(ginLogWriter, os.Stdout)
	return
}
