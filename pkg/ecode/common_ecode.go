package ecode

// All common ecode
var (
	OK = New(200, "success") //请求成功

	ServerErr      = New(500, "server internal error")      //服务器内部错误
	ParamInvaidErr = New(101, "parameter is invalid")       //请求参数错误
	TokenInvidErr  = New(105, "The param token is invalid") //Token校验失败
	PermissionErr  = New(109, "premission denied")          //授权限制
	ForbiddenErr   = New(403, "no privilege")               //禁止进入
)
