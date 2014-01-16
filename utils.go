// utils project utils.go
package utils

import (
	"container/list"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/coffeehc/logger"
)

/*
	字符串list转数据
*/
func ListToStringArray(data *list.List) []string {
	e := data.Front()
	r := make([]string, data.Len())
	length := len(r)
	if length > 0 {
		r[0], _ = e.Value.(string)
		for i := 1; i < length; i++ {
			e = e.Next()
			if e != nil {
				r[i], _ = e.Value.(string)
			} else {
				return r
			}
		}
	}
	return r
}

/*
字符串截取
*/
func SubString(str string, begin, length int) string {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

/*
	计算校验集
*/
func CheckSum(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

/*
	获取文件清单
*/
func FileList(path string) []string {
	list := make([]string, 0)
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		list = append(list, path)
		return nil
	})
	return list
}

/*
	程序暂停n秒,一般测试的时候用得比较多
*/
func WaitTimeOut(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}

/*
	wait,一般是可执行函数的最后用于阻止程序退出
*/
func WaitStop(callBack func()) {
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigChan
	if callBack != nil {
		callBack()
	}
	time.Sleep(time.Second)
}

func StartService(startFunc func(), callBack func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("系统异常退出,原因:%s", err)
			if callBack != nil {
				callBack()
			}
			time.Sleep(time.Second)
		}
	}()
	startFunc()
	WaitStop(callBack)
}
func StartServiceWithArgs(startFunc func(args []string), callBack func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("系统异常退出,原因:%s", err)
			if callBack != nil {
				callBack()
			}
			time.Sleep(time.Second)
		}
	}()
	startFunc(os.Args)
	WaitStop(callBack)
}
