package sql

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

// func TestParseAddrDSN(t *testing.T) {
// 	t.Run("test parse addr dsn", func(t *testing.T) {
// 		addr := parseDSNAddr("test:test@tcp(172.16.0.148:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8")
// 		assert.Equal(t, "172.16.0.148:3306", addr)
// 	})
// 	t.Run("test password has @", func(t *testing.T) {
// 		addr := parseDSNAddr("root:root@dev@tcp(1.2.3.4:3306)/abc?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
// 		assert.Equal(t, "1.2.3.4:3306", addr)
// 	})
// }

// type testConfig struct {
// 	DSN          string        `yaml:"DSN"`          // write data source name.
// 	ReadDSN      []string      `yaml:"ReadDSN"`      // read data source name.
// 	Active       int           `yaml:"Active"`       // pool
// 	Idle         int           `yaml:"Idle"`         // pool
// 	IdleTimeout  time.Duration `yaml:"IdleTimeout"`  // connect max life time.
// 	QueryTimeout time.Duration `yaml:"QueryTimeout"` // query sql timeout
// 	ExecTimeout  time.Duration `yaml:"ExecTimeout"`  // execute sql timeout
// 	TranTimeout  time.Duration `yaml:"TranTimeout"`  // transaction sql timeout
// }

// func TestPing(t *testing.T) {
// 	log.Init(nil)
// 	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
// 	var c Config
// 	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
// 	fmt.Println(c)
// 	db := NewMySQL(&c)
// 	if db == nil {
// 		log.Warn("error")
// 	}

// 	err := db.Ping(context.Background())
// 	if err != nil {
// 		log.Warn(err.Error())
// 	}
// }

type userInfo struct {
	Model
	ID   int    `orm:"id"`
	Name string `orm:"name"`
	Age  int    `orm:"age"`
}

func (u *userInfo) Insert(name string, age int) (err error) {
	_, err = u.Exec(context.Background(), "insert into user_info set name = ?,age = ?", name, age)
	return
}

func (u *userInfo) QueryUserByID(id int) error {
	// rs := u.QueryRow(context.Background(), "select id,name,age from user_info where id = ?", id)
	// if nil != rs {
	// 	fmt.Println(rs)
	// 	e := rs.Scan(&u.ID, &u.Name, &u.Age)
	// 	if e != nil {
	// 		fmt.Println(e.Error())
	// 	}
	// }
	rs, err := u.Query(context.Background(), "select id,name,age from user_info where name = ?", "sun")
	if nil != err {
		log.Error(err.Error())
		panic(err)
	}
	for rs.Next() {
		if err := rs.Scan(&u.ID, &u.Name, &u.Age); nil != err {
			panic(err)
		}
		log.Debug("hhh:%v", u)
	}
	return errors.New("cannot find row")
}

func TestInsert(t *testing.T) {
	log.Init(nil)
	peach.Init(peach.PeachDriverApollo, []string{"zdao_backend.sky_blue", "zdao_backend.common"})
	var c Config
	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
	// fmt.Println(c)
	db := NewMySQL(&c)
	if db == nil {
		log.Warn("error")
	}
	user := userInfo{
		Model: NewModel(db),
	}

	err := user.Begin(context.Background())
	if nil != err {
		panic(err)
	}
	user.Insert("sun", 3)
	// user.QueryUserByID(142)
	fmt.Println(user)
	user.Commit()
	// err := db.Ping(context.Background())
	// if err != nil {
	// 	log.Warn("ping error")
	// }

	// _, err = db.Exec(context.Background(), "insert into user_info set name = ?,age = ?", "name", 423)
	// if nil != err {
	// 	log.Error(err.Error())
	// }
}

type user struct {
	ID   int    `orm:"id"`
	Name string `orm:"name"`
	AgeH int    `orm:"age"`
}

func TestQuery(t *testing.T) {
	// log.Init(nil)
	// c := &Config{
	// 	DSN:          "test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8",
	// 	ReadDSN:      []string{"test:test@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"},
	// 	Active:       10,
	// 	Idle:         10,
	// 	IdleTimeout:  100 * time.Second,
	// 	QueryTimeout: 1 * time.Second,
	// 	ExecTimeout:  1 * time.Second,
	// 	TranTimeout:  1 * time.Second,
	// }
	// db := NewMySQL(c)
	// if db == nil {
	// 	log.Warn("error")
	// }

	// err := db.Ping(context.Background())
	// if err != nil {
	// 	log.Warn(err.Error())
	// }

	// rs, err := db.Query(context.Background(), "select * from user where name = ?", "sun")
	// if nil != err {
	// 	log.Error(err.Error())
	// 	panic(err)
	// }
	// u := user{}
	// for rs.Next() {
	// 	if err := rs.Bind(&u); nil != err {
	// 		panic(err)
	// 	}
	// 	log.Debug("%v", u)
	// }
}
