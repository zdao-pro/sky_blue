package redis

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/common/pool"
)

func TestRedis_Pipeline(t *testing.T) {
	conf := &Config{
		Name:         "test",
		Proto:        "tcp",
		Addr:         testRedisAddr,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	conf.Config = &pool.Config{
		Active:      10,
		Idle:        2,
		IdleTimeout: 90 * time.Second,
	}

	r := NewRedis(conf)
	r.Do(context.TODO(), 0, "FLUSHDB")

	p := r.Pipeline()

	for _, cmd := range testCommands {
		p.Send(cmd.args[0].(string), cmd.args[1:]...)
	}

	replies, err := p.Exec(context.TODO())

	i := 0
	for replies.Next() {
		cmd := testCommands[i]
		actual, err := replies.Scan()
		if err != nil {
			t.Fatalf("Receive(%v) returned error %v", cmd.args, err)
		}
		if !reflect.DeepEqual(actual, cmd.expected) {
			t.Errorf("Receive(%v) = %v, want %v", cmd.args, actual, cmd.expected)
		}
		i++
	}
	err = r.Close()
	if err != nil {
		t.Errorf("Close() error %v", err)
	}
}

func ExamplePipeliner() {
	r := NewRedis(testConfig)
	defer r.Close()

	pip := r.Pipeline()
	pip.Send("SET", "hello", "world")
	pip.Send("GET", "hello")
	replies, err := pip.Exec(context.TODO())
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	for replies.Next() {
		s, err := String(replies.Scan())
		if err != nil {
			fmt.Printf("err %#v\n", err)
		}
		fmt.Printf("%#v\n", s)
	}
	// Output:
	// "OK"
	// "world"
}

func BenchmarkRedisPipelineExec(b *testing.B) {
	r := NewRedis(testConfig)
	defer r.Close()

	r.Do(context.TODO(), 0, "SET", "abcde", "fghiasdfasdf")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := r.Pipeline()
		p.Send("GET", "abcde")
		_, err := p.Exec(context.TODO())
		if err != nil {
			b.Fatal(err)
		}
	}
}
