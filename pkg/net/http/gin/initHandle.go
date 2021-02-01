package gin

//GetInitHandle ..
func GetInitHandle() HandlerFunc {
	return func(c *Context) {
		//init the Context
		// c.Context = context.Background()
	}
}
