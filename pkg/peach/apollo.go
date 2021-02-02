package peach

import (
	"errors"
	"os"
	"strconv"

	apollo "github.com/philchia/agollo/v4"
)

var (
	apolloConfig *Config
)

var nameSpaceList []string

type appoloDriver struct {
}

func (ad *appoloDriver) New(args ...interface{}) (Client, error) {
	nameSpaceArr, ok := args[0].([]string)
	if !ok {
		return nil, errors.New("missing nameSpace list")
	}

	nameSpaceList = nameSpaceArr

	cacheDir := os.Getenv("APOLLO_CACHE_DIR")
	if "" == cacheDir {
		cacheDir = "./"
	}

	client := apollo.NewClient(&apollo.Conf{
		AppID:          os.Getenv("APOLLO_APP_ID"),
		NameSpaceNames: nameSpaceList, // these namespaces will be subscribed at init
		MetaAddr:       os.Getenv("APOLLO_META_ADDR"),
		CacheDir:       cacheDir,
	})

	err := client.Start()
	if err != nil {
		panic(err)
	}

	apolloClientObj := ApoClient{}
	apolloClientObj.client = client
	apolloClientObj.values = new(Map)

	v, err := apolloClientObj.loadValues(nameSpaceList)
	if err != nil {
		panic(err)
	}
	apolloClientObj.values.Store(v)

	// fmt.Println(apolloClientObj.values.Keys())

	return apolloClientObj, nil
}

//Config apollo config struct
type Config struct {
	configServerURL string
	appid           string
	nameSpace       string
}

//ApoClient 实例化的apollo配置中心
type ApoClient struct {
	client apollo.Client
	values *Map
}

//Get 获取NameSpace下key的值
func (ac ApoClient) Get(key string) *Value {
	return ac.values.Get(key)
}

// loadValues load values from apollo namespaces to values
func (ac ApoClient) loadValues(nameSpaceList []string) (values map[string]*Value, err error) {
	values = make(map[string]*Value)
	for _, nameSpace := range nameSpaceList {
		keys := ac.client.GetAllKeys(apollo.WithNamespace(nameSpace))
		for _, k := range keys {
			if values[k], err = ac.loadValue(k, nameSpace); err != nil {
				return
			}
		}
	}
	return
}

// loadValue load value from apollo namespace content to value
func (ac ApoClient) loadValue(key, nameSpace string) (*Value, error) {
	content := ac.client.GetString(key, apollo.WithNamespace(nameSpace))
	i, err := strconv.ParseInt(content, 0, 64)
	if err == nil {
		return NewValue(i, content), nil
	}
	f, err := strconv.ParseFloat(content, 64)
	if err == nil {
		return NewValue(f, content), nil
	}
	b, err := strconv.ParseBool(content)
	if err == nil {
		return NewValue(b, content), nil
	}
	return NewValue(content, content), nil
}
