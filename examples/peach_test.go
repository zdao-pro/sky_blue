package peach

import (
	"fmt"
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

type jsonData struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func TestString(t *testing.T) {
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	a, _ := peach.Get("msyql_test.yaml").String()
	fmt.Println(a)
}

func TestJson(t *testing.T) {
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	var j jsonData
	err := peach.Get("json.json").UnmarshalJSON(&j)
	if nil != err {
		panic(err)
	}
	fmt.Println(j)
}

func TestInt(t *testing.T) {
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	a, _ := peach.Get("test_int").Int()
	fmt.Println(a)
}
