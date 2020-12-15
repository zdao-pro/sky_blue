package sql

import (
	"context"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"

	"github.com/stretchr/testify/assert"
)

func TestParseAddrDSN(t *testing.T) {
	t.Run("test parse addr dsn", func(t *testing.T) {
		addr := parseDSNAddr("test:test@tcp(172.16.0.148:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8")
		assert.Equal(t, "172.16.0.148:3306", addr)
	})
	t.Run("test password has @", func(t *testing.T) {
		addr := parseDSNAddr("root:root@dev@tcp(1.2.3.4:3306)/abc?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
		assert.Equal(t, "1.2.3.4:3306", addr)
	})
}

func TestPing(t *testing.T) {
	log.Init(nil)
	c := &Config{
		DSN:          "test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8",
		ReadDSN:      []string{"test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"},
		Active:       10,
		Idle:         10,
		IdleTimeout:  100 * time.Second,
		QueryTimeout: 1 * time.Second,
		ExecTimeout:  1 * time.Second,
		TranTimeout:  1 * time.Second,
	}
	db := NewMySQL(c)
	if db == nil {
		log.Warn("error")
	}

	err := db.Ping(context.Background())
	if err != nil {
		log.Warn(err.Error())
	}
}

func TestInsert(t *testing.T) {
	log.Init(nil)
	c := &Config{
		DSN:          "test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8",
		ReadDSN:      []string{"test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"},
		Active:       10,
		Idle:         10,
		IdleTimeout:  100 * time.Second,
		QueryTimeout: 1 * time.Second,
		ExecTimeout:  1 * time.Second,
		TranTimeout:  1 * time.Second,
	}
	db := NewMySQL(c)
	if db == nil {
		log.Warn("error")
	}

	err := db.Ping(context.Background())
	if err != nil {
		log.Warn("ping error")
	}

	_, err = db.Exec(context.Background(), "insert into user_info set name = ?,age = ?", "name", 23)
	if nil != err {
		log.Error(err.Error())
	}
}

type user struct {
	id   int
	name string
	age  int
}

func TestQuery(t *testing.T) {
	log.Init(nil)
	c := &Config{
		DSN:          "test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8",
		ReadDSN:      []string{"test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"},
		Active:       10,
		Idle:         10,
		IdleTimeout:  100 * time.Second,
		QueryTimeout: 1 * time.Second,
		ExecTimeout:  1 * time.Second,
		TranTimeout:  1 * time.Second,
	}
	db := NewMySQL(c)
	if db == nil {
		log.Warn("error")
	}

	err := db.Ping(context.Background())
	if err != nil {
		log.Warn(err.Error())
	}

	rs, err := db.Query(context.Background(), "select id,name,age from user_info where name = ?", "name")
	if nil != err {
		log.Error(err.Error())
		panic(err)
	}
	u := user{}
	for rs.Next() {
		if err := rs.Scan(&u.id, &u.name, &u.age); nil != err {
			panic(err)
		}
		log.Debug("%v", u)
	}
}
