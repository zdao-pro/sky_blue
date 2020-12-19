# env SDK

## 服务环境配置
为了更好地实现代码一致性，区分开发，测试，生成环境。
我们引入加入DEV(开发环境) UAT(测试环境) PRO(开发环境) ，通过环境变量实现 
。  

### 环境变量列表

```
	export DEPLOYENV=dev
	export HOSTNAME=user
	export APPID=app
```

| 名称 | 说明 |
|:------|:------|
| DEPLOYENV | dev:开发环境 pre:预发布 online:测试 |
| HOSTNAME | 服务名 |
| APPID | 项目ID |

### 配置
具体见Config结构体,通过Init函数实现初始化

### example
```go
import (
	"fmt"
	"github.com/zdao-pro/sky_blue/pkg/env"
)
// main.go
func main() {
    a := env.GetAppID()
	fmt.Println(a)

	e := env.GetEnv()
	fmt.Println(e)

	h := env.GetHostname()
	fmt.Println(h)
}
```
