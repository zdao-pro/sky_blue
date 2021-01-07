# ecode

## 背景
错误码一般被用来进行异常传递，且需要具有携带`message`文案信息的能力。

## 错误码之Codes

在`sky_blue`里，错误码被设计成`Codes`接口，声明如下[代码位置](https://github.com/zdao-pro/sky_blue/tree/main/pkg/ecode/ecode.go)：

```go
// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
	//Detail get error detail,it may be nil.
	Details() []interface{}
}

// A Code is an int error code spec.
type Code int
```

可以看到该接口一共有四个方法，且`type Code int`结构体实现了该接口。

### 注册message

一个`Code`错误码可以对应一个`message`，默认实现会从全局变量`_messages`中获取，业务可以将自定义`Code`对应的`message`通过调用`Register`方法的方式传递进去，如：

```go
cms := map[int]string{
    0: "success",
    101: "param is error",
}
ecode.Register(cms)

fmt.Println(ecode.OK.Message()) // 输出：很好很强大！
```

注意：`map[int]string`类型并不是绝对，比如有业务要支持多语言的场景就可以扩展为类似`map[int]LangStruct`的结构，因为全局变量`_messages`是`atomic.Value`类型，只需要修改对应的`Message`方法实现即可。

### Details

`Details`接口为`gRPC`预留

## 转换为ecode

错误码转换有以下两种情况：
1. 因为框架传递错误是靠`ecode`错误码，比如bm框架返回的`code`字段默认就是数字，那么客户端接收到如`{"code":404}`的话，可以使用`ec := ecode.Int(404)`或`ec := ecode.String("404")`来进行转换。
2. 在项目中`dao`层返回一个错误码，往往返回参数类型建议为`error`而不是`ecode.Codes`，因为`error`更通用，那么上层`service`就可以使用`ec := ecode.Cause(err)`进行转换。

## 判断

错误码判断是否相等：
1. `ecode`与`ecode`判断使用：`ecode.Equal(ec1, ec2)`
2. `ecode`与`error`判断使用：`ecode.EqualError(ec, err)`


```go
package ecode

import (
    "github.com/zdao-pro/sky_blue/pkg/ecode"
)

var _ ecode.Codes

// UserErrCode
var (
    UserNotLogin = ecode.New(144, 'user no login');
)
```
