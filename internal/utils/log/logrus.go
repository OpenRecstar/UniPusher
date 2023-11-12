package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Log 日志实例
var Log = logrus.New()

// InitLogger 初始化日志配置
func InitLogger() {
	// 设置日志输出为标准输出
	Log.Out = os.Stdout

	// 设置日志级别
	Log.Level = logrus.InfoLevel

	// 设置日志格式为JSON
	Log.Formatter = &logrus.JSONFormatter{}
}
