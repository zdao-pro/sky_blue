package request

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	yaml "gopkg.in/yaml.v2"
)

// KeyNamed key naming to lower case.
func KeyNamed(key string) string {
	return strings.ToLower(key)
}

//Map ..
type Map struct {
	values atomic.Value
}

func init() {
	apiServerListStr := os.Getenv("API_SERVER_LIST")
	apiServerList := strings.Split(apiServerListStr, ",")
	if 0 < len(apiServerList) {
		f := make(map[string]Upstream)
		u := Upstream{
			Server: apiServerList,
		}
		f["api_server"] = u
		src, ok := UpstreamMap.values.Load().(map[string]Upstream)
		if ok {
			for k, v := range src {
				src[k] = v
			}

			for k, v := range f {
				src[k] = v
			}
			UpstreamMap.Store(src)
			return
		}
		UpstreamMap.Store(f)
	}
}

var (
	//UpstreamMap ..
	UpstreamMap Map
	//ErrNoKey ..
	ErrNoKey = fmt.Errorf("the server is not exist")
	//ErrNoMap ..
	ErrNoMap = fmt.Errorf("the UpstreamMap is null")
)

//InitUpstream ..
func InitUpstream(str ...string) {
	for _, s := range str {
		UpstreamMap.Add(s)
	}
}

//Add ..
func (m *Map) Add(s string) {
	f := make(map[string]Upstream)
	err := yaml.Unmarshal([]byte(s), &f)
	if err != nil {
		panic(err)
	}
	src, ok := m.values.Load().(map[string]Upstream)
	if ok {
		for k, v := range src {
			src[k] = v
		}
		m.Store(src)
		return
	}
	m.Store(f)
}

//Store ..
func (m *Map) Store(values map[string]Upstream) {
	m.values.Store(values)
	// fmt.Println(m.values)
}

//Get ..
func (m *Map) Get(key string) (*Upstream, error) {
	src, ok := m.values.Load().(map[string]Upstream)
	if ok {
		v, o := src[key]
		if o {
			return &v, nil
		}
		return nil, ErrNoKey
	}
	return nil, ErrNoMap
}
