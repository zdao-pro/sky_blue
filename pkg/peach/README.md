# peach SDK

## 找到配置中心
为了更好地管理配置，尽可能避免由修改配置带来的失误。
我们引入配置中心模块化思想，可以引入apollo等远程配置中心。
为了区分环境,我们在apollo加入DEV(开发环境) UAT(测试环境) PRO(开发环境)  
。  
## 使用方式

peach 是一个config SDK客户端，包括了remote、file(开发中)几个抽象功能，方便使用本地文件或者远程配置中心，并且集成了对象自动reload功能（开发中）。

### 远程配置中心
需要配合引入环境变量
例如apollo:
```shell
    export APOLLO_META_ADDR=118.178.140.41:58079
    export APOLLO_APP_ID=backend_server
```  

| 名称 | 说明 |
|:------|:------|
| APOLLO_META_ADDR | apollo配置中心地址 |
| APOLLO_APP_ID | apollo的appid |

### 指定配置文件路径：
```shell
    export CONFIG_FILE_PATH=/usr/local
```

### example apollo
```go
import (
	"fmt"

    "github.com/zdao-pro/sky_blue/pkg/peach"
    // 使用apollo需要引入apollo包,实现自动注册驱动
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)
// main.go
func main() {
    // 初始化配置中心
    /*
       @param:
         peach.PeachDriverApollo : 配置驱动名,
         zdao_backend.sky_blue : apollo namespace,apollo需要指定namespace
    */
    peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
    // Get用来获取key的值(value)
	a, _ := peach.Get("test.yaml").String()
	fmt.Println(a)
}
```
