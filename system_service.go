// service
package utils

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/coffeehc/logger"
	"github.com/kardianos/service"
)

type Service interface {
	Run() error
	Start() error
	Stop() error
}

type serviceOwner struct {
	service service.Service
}

func (this *serviceOwner) Run() error {
	return this.service.Run()
}

func (this *serviceOwner) Start() error {
	return this.service.Start()
}

func (this *serviceOwner) Stop() error {
	return this.service.Stop()
}

type serviceWarp struct {
	onstart func()
	onstop  func()
}

func (warp *serviceWarp) Start(s service.Service) error {
	logger.Info("service starting...")
	if warp.onstart == nil {
		return errors.New(logger.Error("没有定义启动方法"))
	}
	go func() {
		defer func() {
			if ok := recover(); ok != nil {
				logger.Error("出现错误:%s,3秒后重新启动")
				time.Sleep(3 * time.Second)
				go s.Restart()
			}
		}()
		warp.onstart()
	}()
	logger.Info("service start ok")
	return nil
}
func (warp *serviceWarp) Stop(s service.Service) error {
	logger.Info("service ending...")
	if warp.onstop != nil {
		warp.onstop()
	} else {
		logger.Warn("没有定义结束的处理方法")
	}
	logger.Info("service end ok")
	return nil
}

func CreatService(name, displayName, desc string, onstart func(), onstop func()) Service {
	config := &service.Config{
		Name:             name,
		DisplayName:      displayName,
		Description:      desc,
		WorkingDirectory: GetAppDir(),
		Arguments:        []string{"run"},
	}
	warp := new(serviceWarp)
	warp.onstart = onstart
	warp.onstop = onstop
	s, err := service.New(warp, config)
	if err != nil {
		log.Error("创建服务[%s]出现错误,%s", name, err)
		return nil
	}
	return &serviceOwner{service: s}
}

func ProcessServiceWithFlag(service_ Service) {
	if !flag.Parsed() {
		flag.Parse()
	}
	var s service.Service = nil
	if ower, ok := service_.(*serviceOwner); ok {
		s = ower.service
	} else {
		logger.Error("没有内置Service对象")
		return
	}
	args := flag.Args()
	if len(args) > 0 {
		var err error
		verb := args[0]
		switch verb {
		case "install":
			fmt.Println("安装")
			err = s.Install()
			if err != nil {
				logger.Error("Failed to install: %s\n", err)
				return
			}
			logger.Info("Service \"%s\" installed.\n", s.String())
			break
		case "uninstall":
			err = s.Uninstall()
			if err != nil {
				logger.Error("Failed to Uninstall: %s\n", err)
				return
			}
			logger.Info("Service \"%s\" Uninstall.\n", s.String())
			break
		case "run":
			err = s.Run()
			if err != nil {
				logger.Error("Failed to run: %s\n", err)
				return
			}
			logger.Info("Service \"%s\" run.\n", s.String())
			break
		case "start":
			err = s.Start()
			if err != nil {
				logger.Error("Failed to start: %s\n", err)
				return
			}
			logger.Info("Service \"%s\" started.\n", s.String())
			break
		case "stop":
			err = s.Stop()
			if err != nil {
				logger.Error("Failed to stop: %s\n", err)
				return
			}
			logger.Info("Service \"%s\" stopped.\n", s.String())
			break
		default:
			logger.Error("不能识别的参数:%#s", args)
		}
		time.Sleep(time.Second)
		return
	}
}
