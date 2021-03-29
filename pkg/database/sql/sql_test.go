package sql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
)

type Misc struct {
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

type UserInfo struct {
	*Model
	ID         int16     `json:"id"`
	Name       string    `json:"name,string"`
	Age        int       `json:"age"`
	Status     uint8     `json:"status"`
	RegistTime time.Time `json:"regist_time"`
	CreateTime time.Time `json:"create_time" time_format:"unixmilli"`
	MiscInfo   Misc      `json:"misc"`
}

func (u *UserInfo) Insert(name string, age int) (err error) {
	_, err = u.Exec(context.Background(), "insert into user_info set name = ?,age = ?", name, age)
	return
}

func (u *UserInfo) QueryUserByID(id int) error {
	err := u.Select(context.Background(), u, "select id,name,age,regist_time,status,create_time,misc from user_info where id = ?", id)
	if nil != err {
		return err
	}
	fmt.Println(u)
	return nil
}

func (u *UserInfo) QueryUserByName(name string) (*[]UserInfo, error) {
	list := make([]UserInfo, 0)
	err := u.Select(context.Background(), &list, "select id,name,age,regist_time,status,create_time,misc from user_info where name = ?", name)
	if nil != err {
		return nil, err
	}
	// fmt.Println(list)
	return &list, nil
}

type Man struct {
	Id   int
	Name string
}

func TestInsert(t *testing.T) {
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
	user := UserInfo{
		Model: NewModel(db),
	}
	// user.Begin(context.Background())
	// user.Insert("test", 50)
	// user.Rollback()
	user.QueryUserByID(145)
	fmt.Println(user)
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	AgeH int    `json:"age"`
}
