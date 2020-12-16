package apollo

import (
	"errors"
	"fmt"
	"os"

	"github.com/zdao-pro/sky_blue/pkg/peach"

	"github.com/shima-park/agollo"
)

var (
	apolloConfig *Config
)

func init() {
	/*
		export APOLLO_META_ADDR=118.178.140.41:58079
		export APOLLO_APP_ID=backend_server
	*/
	peach.RegistDriver("apollo", &appoloDriver{})
}

type appoloDriver struct {
}

func (ad *appoloDriver) New(args ...interface{}) (peach.Client, error) {
	configServerURL := os.Getenv("APOLLO_META_ADDR")
	if "" == configServerURL {
		return nil, errors.New("missing APOLLO_META_ADDR")
	}
	appid := os.Getenv("APOLLO_APP_ID")
	if "" == appid {
		return nil, errors.New("missing APOLLO_APP_ID")
	}
	nameSpace, ok := args[0].(string)
	if !ok {
		return nil, errors.New("missing nameSpace")
	}
	apolloConfig = &Config{
		configServerURL: configServerURL,
		appid:           appid,
		nameSpace:       nameSpace,
	}
	apolloClientObj, err := newApolloClient(apolloConfig)
	if nil != err {
		return nil, fmt.Errorf("paladin: unknown appoloDriver (forgotten register?)")
	}

	return apolloClientObj, nil
}

//Config apollo config struct
type Config struct {
	configServerURL string
	appid           string
	nameSpace       string
}

//newApolloClient 返回一个apollo配置中心
func newApolloClient(c *Config) (peach.Client, error) {

	a, err := agollo.New(c.configServerURL, c.appid, agollo.AutoFetchOnCacheMiss())
	if err != nil {
		return nil, err
	}

	errorCh := a.Start()

	if errorCh != nil {
		// fmt.Println("Start failed.......")
	}

	if nil != err {
		panic(err)
	}
	apoloClientInstance := &ApoClient{}
	apoloClientInstance.SetApollo(a)

	return apoloClientInstance, nil
}

//ApoClient 实例化的apollo配置中心
type ApoClient struct {
	a agollo.Agollo
}

//SetApollo 设置apollo客户端
func (ac *ApoClient) SetApollo(ap agollo.Agollo) {
	ac.a = ap
}

//Get 获取NameSpace下key的值
func (ac *ApoClient) Get(k string) *peach.Value {
	v := ac.a.Get(k, agollo.WithNamespace(apolloConfig.nameSpace))
	return peach.NewValue(v, v)
}

//GetKey 获取NameSpace下key的值
func (ac *ApoClient) GetKey(key string) string {
	k := ac.a.Get(key, agollo.WithNamespace(apolloConfig.nameSpace))
	fmt.Println(k)
	return k
}
