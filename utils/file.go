package utils

//文件操作包
import (
	"io"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// 动态化改变文件流
func RotationFile(logFileName string) (w io.Writer, err error) {
	fileName := path.Join(ConfigData.Path, logFileName)
	logf, err := rotatelogs.New(
		fileName,
		rotatelogs.WithRotationTime(time.Duration(ConfigData.RotationTime)),
		rotatelogs.WithRotationSize(ConfigData.RotationSize*1024),
	)
	if err != nil {
		return nil, err
	}
	return io.MultiWriter(logf), err
}
