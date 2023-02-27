package utils

import (
	"go.uber.org/zap"
)

func GetLoger() *zap.Logger {
	return zap.L()
}

func GetSugarLogerAndLoger() (*zap.SugaredLogger, *zap.Logger) {
	loger := GetLoger()
	return loger.Sugar(), loger
}
