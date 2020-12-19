package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

type testConfig struct {
	// Active number of items allocated by the pool at a given time.
	// When zero, there is no limit on the number of items in the pool.
	Active int `yaml:"active"`
	// Idle number of idle items in the pool.
	Idle int `yaml:"idle"`
	// Close items after remaining item for this duration. If the value
	// is zero, then item items are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration `yaml:"idleTimeout"`
	// If WaitTimeout is set and the pool is at the Active limit, then Get() waits WatiTimeout
	// until a item to be returned to the pool before returning.
	WaitTimeout time.Duration `yaml:"waitTimeout"`
	// If WaitTimeout is not set, then Wait effects.
	// if Wait is set true, then wait until ctx timeout, or default flase and return directly.
	Wait         bool          `yaml:"wait"`
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
	var c redis.NewConfig
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
	r := redis.NewRedisClient(&c)
	defer r.Close()
	r.Do(context.Background(), 4, "SET", "a", "f")
}
