package request

import (
	"fmt"
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
