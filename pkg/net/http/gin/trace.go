package gin

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zdao-pro/sky_blue/pkg/net/trace"
)

const (
	_TraceName = "gin_trace"
)

//ginTrace ..
type ginTrace struct {
	traceID string
}

//TraceID ..
func (g *ginTrace) TraceID() string {
	return g.traceID
}

//SetTraceID ..
func (g *ginTrace) SetTraceID(t string) {
	g.traceID = t
}

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
		t := &ginTrace{}
		t.SetTraceID(traceID)
		fmt.Println("t:", t.TraceID())
		//new c.Context
		c.Context = trace.NewContext(c.Context, t)
	}
}
