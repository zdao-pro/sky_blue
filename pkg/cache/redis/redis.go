// Copyright 2012 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package redis

import (
	"context"

	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/zdao-pro/sky_blue/pkg/common/pool"
)

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

// Config client settings.
type Config struct {
	*pool.Config

	Name         string // redis name, for trace
	Proto        string
	Addr         string
	Auth         string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SlowLog      time.Duration
}

// NewConfig ..
type NewConfig struct {
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

// Redis ..
type Redis struct {
	pool *Pool
	conf *Config
}

// NewRedis ..
func NewRedis(c *Config, options ...DialOption) *Redis {
	return &Redis{
		pool: NewPool(c, options...),
		conf: c,
	}
}

//NewRedisClient ..
func NewRedisClient(c *NewConfig, options ...DialOption) *Redis {
	poolConfig := &pool.Config{
		Active:      c.Active,
		Idle:        c.Idle,
		IdleTimeout: c.IdleTimeout,
		WaitTimeout: c.WaitTimeout,
	}
	redisConfig := Config{
		Config:       poolConfig,
		Name:         c.Name,
		Proto:        c.Proto,
		Addr:         c.Addr,
		Auth:         c.Auth,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		SlowLog:      c.SlowLog,
	}
	redisConfig.Wait = c.Wait
	return &Redis{
		pool: NewPool(&redisConfig, options...),
		conf: &redisConfig,
	}
}

// Do gets a new conn from pool, then execute Do with this conn, finally close this conn.
// ATTENTION: Don't use this method with transaction command like MULTI etc. Because every Do will close conn automatically, use r.Conn to get a raw conn for this situation.
func (r *Redis) Do(ctx context.Context, db int, commandName string, args ...interface{}) (reply interface{}, err error) {
	//trace
	s := opentracing.SpanFromContext(ctx)
	if nil != s {
		span2 := opentracing.StartSpan(commandName, opentracing.ChildOf(s.Context()))
		span2.SetTag("commandName", commandName)
		span2.SetTag("db", db)
		span2.SetTag("args", args)
		defer span2.Finish()
	}
	conn := r.pool.Get(ctx)
	defer conn.Close()
	if _, err := conn.Do("SELECT", db); err != nil {
		conn.Close()
		return nil, errors.WithStack(err)
	}
	reply, err = conn.Do(commandName, args...)
	return
}

// Close closes connection pool
func (r *Redis) Close() error {
	return r.pool.Close()
}

// Conn direct gets a connection
func (r *Redis) Conn(ctx context.Context) Conn {
	return r.pool.Get(ctx)
}

// Pipeline ..
func (r *Redis) Pipeline() (p Pipeliner) {
	return &pipeliner{
		pool: r.pool,
	}
}
