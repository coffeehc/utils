// service
package utils

import (
	"flag"
	"fmt"
	"time"

	"bitbucket.org/kardianos/service"
)

func StartService(name, displayName, desc string, onstart func(failMsg chan string), onstop func()) {
	s, err := service.NewService(name, displayName, desc)
	if err != nil {
		log.Error("创建服务[%s]出现错误,%s", name, err)
		return
	}
	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	if len(args) > 1 {
		var err error
		verb := args[1]
		switch verb {
		case "install":
			fmt.Println("安装")
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" installed.\n", displayName)
			break
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", displayName)
			break
		case "run":
			stop := make(chan string)
			go func() {
				failMsg := <-stop
				log.Error("启动失败，原因：%s", failMsg)
				s.Stop()
			}()
			onstart(stop)
			break
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", displayName)
			break
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", displayName)
			break
		default:
			fmt.Printf("不能识别的参数:%#s", args)
		}
		time.Sleep(time.Second)
		return
	}
	err = s.Run(func() error {
		stop := make(chan string)
		go func() {
			failMsg := <-stop
			log.Error("启动失败，原因：%s", failMsg)
			s.Stop()
		}()
		go onstart(stop)
		return nil
	}, func() error {
		if onstop != nil {
			onstop()
		}
		log.Info("service 停止")
		return nil
	})
	if err != nil {
		log.Error(err.Error())
	}
}
