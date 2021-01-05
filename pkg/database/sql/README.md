# database-mysql模块

## 找到database-mysql
摘自bilibili karos框架,加入自写的ORM框架
## 使用方式

database-mysql 是一个mysql客户端，包括了基本DML,性能监控,可以使用peach配置SDK加载配置文件。

### example

示例数据表结构：
> misc字段存储的是json字符串
```
+-------------+---------------+------+-----+-------------------+-------------------+
| Field       | Type          | Null | Key | Default           | Extra             |
+-------------+---------------+------+-----+-------------------+-------------------+
| id          | bigint        | NO   | PRI | NULL              | auto_increment    |
| name        | varchar(50)   | NO   | MUL |                   |                   |
| age         | int           | YES  |     | NULL              |                   |
| regist_time | timestamp     | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
| status      | tinyint       | NO   |     | 0                 |                   |
| create_time | bigint        | NO   |     | 0                 |                   |
| misc        | varchar(1024) | NO   |     |                   |                   |
+-------------+---------------+------+-----+-------------------+-------------------+
```

#### 1. 普通使用
```go
import (
	"context"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
    "github.com/zdao-pro/sky_blue/pkg/database/sql"
	"github.com/stretchr/testify/assert"
)

func main() {
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
	db := sql.NewMySQL(c)
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
```

#### 2. 配置文件使用方式
```go
import (
	"context"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/database/sql"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func main() {
	log.Init(nil)
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	var c Config
	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
	fmt.Println(c)
	db := sql.NewMySQL(&c)
	if db == nil {
		log.Warn("error")
	}

	err := db.Ping(context.Background())
	if err != nil {
		log.Warn("ping error")
	}

	_, err = db.Exec(context.Background(), "insert into user_info set name = ?,age = ?", "name", 423)
	if nil != err {
		log.Error(err.Error())
	}
}
```

#### 3. ORM方式(推荐)
```go
import (
	"context"
	"time"
	"error"

    "github.com/zdao-pro/sky_blue/pkg/database/sql"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

// 使用时需在结构体tag指明orm
// 当字段存储的是int类型时间戳，需指明time_format:  unix(秒级时间戳) unixmilli(毫秒级时间戳) unixnano(纳秒级时间戳)，才可以自动解析为time.Time
// 当字段为json字符串,需要在解析结构体tag指明json, 才可以自动解析
// 结构体需要声明*sql.Model,并赋值
// 使用事务时,可以直接使用Begin(context.Background()),Commit(),Rollback(),不需要另外声明
type UserInfo struct {
	*sql.Model
	ID         int16     `orm:"id"`
	Name       string    `orm:"name"`
	Age        int       `orm:"age"`
	Status     uint8     `orm:"status"`
	RegistTime time.Time `orm:"regist_time"`
	CreateTime time.Time `orm:"create_time" time_format:"unixmilli"`
	MiscInfo   Misc      `orm:"misc"`
}

type Misc struct {
	Email string `json:"email"`
	Phone int    `json:"phone"`
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

func main() {
	log.Init(nil)
	peach.Init(peach.PeachDriverApollo, []string{"zdao_backend.sky_blue", "zdao_backend.common"})
	var c Config
	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
	db := sql.NewMySQL(&c)
	if db == nil {
		log.Warn("error")
	}
	user := UserInfo{
		Model: NewModel(db),
	}
	// user.Begin(context.Background())
	// user.Insert("test", 50)
	// user.Rollback()
	// user.Commit()
	user.QueryUserByID(145)
	fmt.Println(user)
}

```
