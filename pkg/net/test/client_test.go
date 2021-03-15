package main

import (
	"context"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/naming/etcd"
	"github.com/zdao-pro/sky_blue/pkg/net/rpc/warden"
	"github.com/zdao-pro/sky_blue/pkg/net/rpc/warden/resolver"
	"google.golang.org/grpc"
)

func main() {
	log.Init(nil)
	resolver.Register(etcd.Builder(nil))
}

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	// 这里使用etcd scheme
	conn, err := client.Dial(context.Background(), "etcd://default/"+AppID)
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return NewDemoClient(conn), nil
}
