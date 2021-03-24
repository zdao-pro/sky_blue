package etcd

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/zdao-pro/sky_blue/pkg/naming"
	"go.etcd.io/etcd/clientv3"
)

func TestRegister(t *testing.T) {
	c := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}
	b, err := New(&c)
	if nil != err {
		fmt.Println(err.Error())
	}
	in := naming.Instance{
		AppID:    "jim",
		Hostname: "hh",
	}
	_, err = b.Register(context.Background(), &in)
	if nil != err {
		fmt.Println(err.Error())
	}
}
