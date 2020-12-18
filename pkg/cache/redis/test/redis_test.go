package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/common/pool"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

type testConfig struct {
	*pool.Config

	Name         string        `yaml:"name"`
	Proto        string        `yaml:"proto"`
	Addr         string        `yaml:"addr"`
	Auth         string        `yaml:"auth"`
	DialTimeout  time.Duration `yaml:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	SlowLog      time.Duration `yaml:"slowLog"`
}

func TestApollo(t *testing.T) {
	log.Init(nil)
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	var c testConfig
	peach.Get("redis_test.yaml").UnmarshalYAML(&c)
	fmt.Println(c)
	// config := &redis.Config{
	// 	Config: &pool.Config{
	// 		Active: 10,
	// 		Idle:   5,
	// 	},
	// 	Name:         "test_get",
	// 	Proto:        "tcp",
	// 	Addr:         "localhost:6379",
	// 	DialTimeout:  time.Second,
	// 	ReadTimeout:  time.Second,
	// 	WriteTimeout: time.Second,
	// }
	// r := redis.NewRedis(config)
	// defer r.Close()
	// r.Do(context.Background(), "SET", "a", "c")
}
