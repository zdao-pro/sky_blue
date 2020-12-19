# cache-redis模块

## 找到cache-redis
摘自bilibili karos框架,对其简单修改
## 使用方式

cache-redis 是一个redis客户端，包括了基本DML,性能监控,可以使用peach配置SDK加载配置文件。

### example 

#### 1. 普通使用
```go
import (
	"context"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
	"github.com/zdao-pro/sky_blue/pkg/common/pool"
)

func main() {
	config := &redis.Config{
		Config: &pool.Config{
			Active: 10, // 连接池活跃数
			Idle:   5, // 连接池空闲数
		},
		Name:         "test_get",
		Proto:        "tcp",
		Addr:         "localhost:6379",
		DialTimeout:  time.Second, // 连接超时
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	db := 1 // 指定数据库
	r := redis.NewRedis(config) // 内部维护一个redis连接池
	defer r.Close()
	r.Do(context.Background(), db, "SET", "a", "b")
}
```

#### 1. 配置文件使用方式
```go
import (
	"context"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
	"github.com/zdao-pro/sky_blue/pkg/common/pool"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func main() {
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	var c redis.NewConfig
	peach.Get("redis_test.yaml").UnmarshalYAML(&c)
	fmt.Println(c)
	r := redis.NewRedisClient(&c)
	db := 2 // 指定数据库
	defer r.Close()
	r.Do(context.Background(), db, "SET", "a", "f")
}
```
