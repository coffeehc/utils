/**
 * Created by coffee on 15/11/15.
 */
package utils
import (
	"os"
	"syscall"
	"time"
	"os/signal"
	"github.com/coffeehc/logger"
)

type serviceWarp struct {
	runFunc  func()
	stopFunc func()
}

func (this *serviceWarp)Run() {
	this.runFunc()
}

func (this *serviceWarp)Stop() {
	this.Stop()
}
func NewService(runFunc func(), stopFunc func()) Service {
	return &serviceWarp{runFunc, stopFunc}
}

func StartService(service Service) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("service crash,cause is %s", err)
		}
		if service.Stop != nil {
			service.Stop()
		}
		time.Sleep(time.Second)
	}()
	if service.Run == nil {
		panic("没有指定Run方法")
	}
	WaitStop()
}

type Service interface {
	Run()
	Stop()
}

/*
	wait,一般是可执行函数的最后用于阻止程序退出
*/
func WaitStop() {
	var sigChan = make(chan os.Signal, 3)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	sig := <-sigChan
	logger.Debug("接收到指令:%s,立即关闭程序", sig)
}