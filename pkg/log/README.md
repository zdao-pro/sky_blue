# log SDK

## 找到日志模块
为了更好地提供日志功能。
为了区分环境,我们引入加入DEV(开发环境) UAT(测试环境) PRO(开发环境)  
。  

### 使用方式
1 日志环境变量,具体见env模块
```
1 对于不同的服务器环境，日志会选择不同的handle:
  例如:
      对于开发环境，会把日志打印到控制台
      对于线上，测试环境，日志会打印到远程日志中心。
```

```shell
    export UDP_LOG_ADDR=118.178.140.41
    export UDP_LOG_PORT=1223
    export SYSLOG_ADDR=118.178.140.41
    export SYSLOG_PORT=1224
```  

| 名称 | 说明 |
|:------|:------|
| UDP_LOG_ADDR | UDP远程日志中心IP地址 |
| UDP_LOG_PORT | UDP远程日志中心port |
| SYSLOG_ADDR | syslog远程日志中心IP地址 |
| SYSLOG_PORT | syslog远程日志中心port |

### 配置
具体见Config结构体,通过Init函数实现初始化

### example
```go
import (
	"github.com/zdao-pro/sky_blue/pkg/log"
)
// main.go
func main() {
    log.Init(nil)
	log.Info("%s", "222322")
}
```
 