package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//日志文件

var loger *zap.Logger

// 日志格式设置
func getEncoder() zapcore.Encoder {
	encodeCofnig := zap.NewProductionEncoderConfig()
	encodeCofnig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeCofnig.TimeKey = "time"
	encodeCofnig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeCofnig)
}

//设置日志写入位置

func getLogWriter(logFileName string) zapcore.WriteSyncer {
	w, _ := RotationFile(logFileName)
	return zapcore.AddSync(w)
}

// 初始化并替代全局loger
func InitLoger(logFileName string) (err error) {
	writeSync := getLogWriter(logFileName)
	encoder := getEncoder()
	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(ConfigData.Log.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSync, l)
	loger = zap.New(core, zap.AddCaller())
	//把loger替换成全局实例，这样在其他文件只要zap.L()就可以使用了
	zap.ReplaceGlobals(loger)
	return
}

func GetLoger() *zap.Logger {
	return zap.L()
}

func GetSugarLogerAndLoger() (*zap.SugaredLogger, *zap.Logger) {
	loger := GetLoger()
	return loger.Sugar(), loger
}
