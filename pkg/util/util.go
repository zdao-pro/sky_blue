package util

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"crypto/md5"

	"github.com/google/uuid"
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

// JoinStringsInASCIIAscend 对map按key字典升序拼接
/**
	* @param sep 拼接字符串
**/
func JoinStringsInASCIIAscend(data map[string]interface{}, sep string) string {
	var keyList []string
	for key := range data {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)
	var strList []string
	for _, key := range keyList {
		strList = append(strList, fmt.Sprintf("%s=%v", key, data[key]))
	}
	return strings.Join(strList, sep)
}

// ParseQuery 解析参数(a=b&c=q)
func ParseQuery(url string) (map[string]string, error) {
	query := make(map[string]string, 0)
	for _, val := range strings.Split(url, "&") {
		v := strings.Split(val, "=")
		if len(v) < 2 {
			return make(map[string]string, 0), errors.New("query parameter error")
		}
		query[v[0]] = v[1]
	}
	return query, nil
}

// GetUUID 获取uuid
func GetUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return id.String()
}

// MD5 返回16进制字符串(32位)
func MD5(src []byte) string {
	return fmt.Sprintf("%x", md5.Sum(src))
}
