package util

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"runtime"
)

var (
	funcRegex = regexp.MustCompile(`^.+[\\/]([\d\w\.\*\(\)_]+)$`)
)

// FuncName get func name.
// 获取函数名(包含文件名前缀)
func FuncName(skip int) (name string) {
	pc := make([]uintptr, 2)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	if nil != f {
		s := f.Name()
		a := funcRegex.FindStringSubmatch(s)
		if 2 == len(a) {
			return a[1]
		}
		return s
	}
	return "unknown:0"
}

//GetLocalAddress ...
func GetLocalAddress() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"
}
