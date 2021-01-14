package gin

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zdao-pro/sky_blue/pkg/net/trace"
)

// Trace is trace middleware
func Trace() HandlerFunc {
	return func(c *Context) {
		s := c.GetHeader("trace_id")
		var traceID string
		if "" != s {
			traceID = s
		} else {
			if uu, err := uuid.NewUUID(); err == nil {
				traceID = uu.String()
			}

		}
		t := trace.NewGinTrace(traceID)
		fmt.Println("t:", t.TraceID())
		//new c.Context
		c.Context = trace.NewContext(c.Context, t)

	}
}
