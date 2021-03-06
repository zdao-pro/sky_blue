package main

import (
	"context"
	"fmt"

	pb "github.com/zdao-pro/sky_blue/examples/proto/testproto"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/naming/etcd"
	"github.com/zdao-pro/sky_blue/pkg/net/rpc/warden"
	"github.com/zdao-pro/sky_blue/pkg/net/rpc/warden/resolver"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

func main() {
	log.Init(nil)
	c := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}
	resolver.Register(etcd.Builder(&c))
	conn, err := NewClient(nil)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		defer conn.Close()
		c := pb.NewGreeterClient(conn)
		resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "sy", Age: 13})
		if err != nil {
			fmt.Println("err:", err.Error())
		} else {
			fmt.Println("res:", *resp)
		}
	}
}

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	client := warden.NewClient(cfg, opts...)
	var AppID string = "jim"
	// 这里使用etcd scheme
	conn, err := client.Dial(context.Background(), "etcd://default/"+AppID)
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return conn, nil
}
