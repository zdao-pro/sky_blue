package redis

import (
	"context"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
	"github.com/zdao-pro/sky_blue/pkg/common/pool"
)

func TestApollo(t *testing.T) {
	config := &redis.Config{
		Config: &pool.Config{
			Active: 10,
			Idle:   5,
		},
		Name:         "test_get",
		Proto:        "tcp",
		Addr:         "localhost:6379",
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	r := redis.NewRedis(config)
	defer r.Close()
	r.Do(context.Background(), "SET", "a", "c")
}
