// utils project utils.go
package utils

import (
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

/*
	获取App执行文件目录
*/
func GetAppDir() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

/*
	wait,一般是可执行函数的最后用于阻止程序退出
*/
func WaitStop(callBack func()) {
	var sigChan = make(chan os.Signal, 3)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigChan
	if callBack != nil {
		callBack()
	}
	time.Sleep(time.Second)
}
