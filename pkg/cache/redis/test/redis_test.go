package redis

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"bufio"

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

func TestIn(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("     len cap   address")
	fmt.Print("111---", len(nums), cap(nums))
	fmt.Printf("    %p\n", nums) //0xc4200181e0
	a := nums[:3]
	fmt.Print("222---", len(a), cap(a))
	fmt.Printf("    %p\n", a) //0xc4200181e0 一样
	//output: 222--- 3 5

	var b = make([]int, 3)
	// b := nums[:3:3]          //第二个冒号 设置cap的
	n := copy(b, nums[:3:3]) //第二个冒号 设置cap的
	fmt.Print("333---", len(b), cap(b))
	fmt.Printf("    %p\n", b) //0xc4200181e0 一样
	//output: 333--- 3 3
	fmt.Println(n, b)
	nums[0] = 55
	a[0] = 9
	fmt.Println(nums, a, b)
}

func TestByte(t *testing.T) {
	b := make([]byte, 3, 5)
	fmt.Println(b, " ptr:", &b[0])
	b = b[:4]
	fmt.Println(b, " ptr:", &b[0])
}

func TestRune(t *testing.T) {
	var str = "hello 你好"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))
}

func TestBufIo(t *testing.T) {
	var str = "irhquihrqr$wriqrqirh$hqrhqu"
	b := bufio.NewReader(strings.NewReader(str))
	s, err := b.ReadSlice(byte('$'))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
	i := bytes.IndexByte([]byte(str), byte('$'))
	fmt.Println(string([]byte(str)[0:i+1]), " i:", i)
}
