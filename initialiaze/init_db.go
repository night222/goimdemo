package initialiaze

import (
	"fmt"
	"goimdemo/common"
	"goimdemo/utils"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化mysql连接
func InitMysql() (err error) {
	//设置把sql日志输出到控制台
	sqlLogWriter, err := utils.RotationFile("/sql_log_%Y-%m-%d.log")
	if err != nil {
		return
	}
	newLOger := logger.New(log.New(sqlLogWriter, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second, //慢sql阈值
		LogLevel:                  logger.Info, //级别
		Colorful:                  false,       //彩色
		IgnoreRecordNotFoundError: false,       //是否显示没有查询到的结果
	})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s%s",
		common.ConfigData.UserName,
		common.ConfigData.Passworld,
		common.ConfigData.DatabaseConfig.Host,
		common.ConfigData.DatabaseConfig.Port,
		common.ConfigData.DatabaseName,
		common.ConfigData.Charset,
		common.ConfigData.ParseTime,
		common.ConfigData.Supplement,
	)
	common.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLOger,
	})
	return
}
