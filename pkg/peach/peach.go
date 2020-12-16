package peach

import (
	"errors"
)

var (
	confPath string //配置文件路径
	//ErrInitConfigException 配置中心初始化异常
	ErrInitConfigException = errors.New("Init Config Center Error")
	//DefaultClient default config client
	DefaultClient Client
)

func init() {

}

//Init 初始化配置中心
func Init(driverName string, args ...interface{}) error {
	driver, err := GetDriver(driverName)
	if err != nil {
		panic(err)
	}

	var ok error
	DefaultClient, ok = driver.New(args...)

	if ok != nil {
		return nil
	}
	return nil
}

// Get return value by key.
func Get(key string) *Value {
	return DefaultClient.Get(key)
}
