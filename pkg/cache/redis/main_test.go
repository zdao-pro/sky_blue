package redis

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/pkg/testing/lich"
	"github.com/zdao-pro/sky_blue/pkg/common/pool"
)

var (
	testRedisAddr string
	testPool      *Pool
	testConfig    *Config
)

func setupTestConfig(addr string) {
	c := getTestConfig(addr)
	c.Config = &pool.Config{
		Active:      20,
		Idle:        2,
		IdleTimeout: 90 * time.Second,
	}
	testConfig = c
}

func getTestConfig(addr string) *Config {
	return &Config{
		Name:         "test",
		Proto:        "tcp",
		Addr:         addr,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
}

func setupTestPool() {
	testPool = NewPool(testConfig)
}

// DialDefaultServer starts the test server if not already started and dials a
// connection to the server.
func DialDefaultServer() (Conn, error) {
	c, err := Dial("tcp", testRedisAddr, DialReadTimeout(1*time.Second), DialWriteTimeout(1*time.Second))
	if err != nil {
		return nil, err
	}
	c.Do("FLUSHDB")
	return c, nil
}

func TestMain(m *testing.M) {
	flag.Set("f", "./test/docker-compose.yaml")
	if err := lich.Setup(); err != nil {
		panic(err)
	}
	defer lich.Teardown()
	testRedisAddr = "localhost:6379"
	setupTestConfig(testRedisAddr)
	setupTestPool()
	ret := m.Run()
	os.Exit(ret)
}
