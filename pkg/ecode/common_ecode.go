package ecode

// All common ecode
var (
	OK = New(200, "success") //请求成功

	ServerErr          = New(500, "server internal error")      //服务器内部错误
	ParamInvaidErr     = New(101, "parameter is invalid")       //请求参数错误
	TokenInvidErr      = New(105, "The param token is invalid") //Token校验失败
	PermissionErr      = New(109, "premission denied")          //授权限制
	ForbiddenErr       = New(403, "no privilege")               //禁止进入
	NotModified        = New(-304, "error")                     // 木有改动
	TemporaryRedirect  = New(-307, "error")                     // 撞车跳转
	RequestErr         = New(-400, "error")                     // 请求错误
	Unauthorized       = New(-401, "error")                     // 未认证
	AccessDenied       = New(-403, "error")                     // 访问权限不足
	NothingFound       = New(-404, "error")                     // 啥都木有
	MethodNotAllowed   = New(-405, "error")                     // 不支持该方法
	Conflict           = New(-409, "error")                     // 冲突
	Canceled           = New(-498, "error")                     // 客户端取消请求
	ServiceUnavailable = New(-503, "error")                     // 过载保护,服务暂不可用
	Deadline           = New(-504, "error")                     // 服务调用超时
	LimitExceed        = New(-509, "error")                     // 超出限制
)
