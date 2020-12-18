# database-mysql模块

## 找到database-mysql
摘自bilibili karos框架,对其简单修改
## 使用方式

database-mysql 是一个mysql客户端，包括了基本DML,性能监控,可以使用peach配置SDK加载配置文件。

### example 

#### 1. 普通使用
```go
import (
	"context"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"

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
```

#### 2. 配置文件使用方式
