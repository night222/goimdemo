package initialiaze

import (
	"goimdemo/common"
	"goimdemo/config"
)

// 初始化配置文件
func InitConfig() (err error) {
	common.ConfigData, err = config.InitConfig()
	return
}
