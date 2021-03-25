package redis

import (
	"context"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
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
	// var c redis.NewConfig
	// peach.Get("db_redis_user_persist.yaml").UnmarshalYAML(&c)
	// fmt.Println(c)
	config := &redis.NewConfig{
		Active:       10,
		Idle:         5,
		Name:         "test_get",
		Proto:        "tcp",
		Addr:         "r-bp1mwqdr5khc6uui7dpd.redis.rds.aliyuncs.com:6379",
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Auth:         "zhaodao_2020",
	}
	r := redis.NewRedisClient(config)
	defer r.Close()
	_, err := r.Do(context.Background(), 4, "SET", "a", "f")
	if err != nil {
		t.Fatalf(err.Error())
	}
}
