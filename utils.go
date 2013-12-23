// utils project utils.go
package utils

import (
	"container/list"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

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

func FileList(path string, subFile bool) []string {
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

func WaitTimeOut(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}

func WaitStop(callBack func()) {
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan
	if callBack != nil {
		callBack()
	}
	time.Sleep(time.Second)
}
