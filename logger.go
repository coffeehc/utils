// logger
package utils

import "github.com/coffeehc/logger"

//日子接口,返回具体的日志内容
type Logger interface {
	Trace(format string, v ...interface{}) string
	Debug(format string, v ...interface{}) string
	Info(format string, v ...interface{}) string
	Warn(format string, v ...interface{}) string
	Error(format string, v ...interface{}) string
}

const (
	LOGGER_TIMEFORMAT_NANOSECOND string = "2006-01-02 15:04:05.999999999"
)

var log Logger = logger.GetLogger()

func SetLogger(logger Logger) {
	log = logger
}
