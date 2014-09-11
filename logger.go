// logger
package utils

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger interface {
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}

const (
	LOGGER_TIMEFORMAT_NANOSECOND string = "2006-01-02 15:04:05.999999999"
)

var log Logger = &utilLogger{}

type utilLogger struct {
}

func getLineInfo() string {
	_, file, line, ok := runtime.Caller(3)
	var lineInfo string = "-:0"
	if ok {
		index := strings.Index(file, "/src/") + 4
		lineInfo = file[index:] + ":" + strconv.Itoa(line)
	}
	return lineInfo
}

func (this utilLogger) Trace(format string, v ...interface{}) {
	fmt.Printf("%s\t%s\t%s\t%s\n", time.Now().Format(LOGGER_TIMEFORMAT_NANOSECOND), "TRACE", getLineInfo(), fmt.Sprintf(format, v...))
}

func (this utilLogger) Debug(format string, v ...interface{}) {
	fmt.Printf("%s\t%s\t%s\t%s\n", time.Now().Format(LOGGER_TIMEFORMAT_NANOSECOND), "DEBUG", getLineInfo(), fmt.Sprintf(format, v...))
}

func (this utilLogger) Info(format string, v ...interface{}) {
	fmt.Printf("%s\t%s\t%s\t%s\n", time.Now().Format(LOGGER_TIMEFORMAT_NANOSECOND), "INFO", getLineInfo(), fmt.Sprintf(format, v...))
}

func (this utilLogger) Warn(format string, v ...interface{}) {
	fmt.Printf("%s\t%s\t%s\t%s\n", time.Now().Format(LOGGER_TIMEFORMAT_NANOSECOND), "WARN", getLineInfo(), fmt.Sprintf(format, v...))
}

func (this utilLogger) Error(format string, v ...interface{}) {
	fmt.Printf("%s\t%s\t%s\t%s\n", time.Now().Format(LOGGER_TIMEFORMAT_NANOSECOND), "ERROR", getLineInfo(), fmt.Sprintf(format, v...))
}

func SetLogger(logger Logger) {
	log = logger
}
