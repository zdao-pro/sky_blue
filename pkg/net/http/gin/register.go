package gin

//RouteRegister 注册route
type RouteRegister interface {
	RegisterDemoBMServer(e *Engine)
}
