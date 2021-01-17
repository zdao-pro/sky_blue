package gin

import (
	"github.com/opentracing/opentracing-go"
)

// Trace is trace middleware
func Trace() HandlerFunc {
	return func(c *Context) {
		// s := c.GetHeader("trace_id")
		// var traceID string
		// if "" != s {
		// 	traceID = s
		// } else {
		// 	if uu, err := uuid.NewUUID(); err == nil {
		// 		traceID = uu.String()
		// 	}

		// }
		// t := trace.NewGinTrace(traceID)
		// fmt.Println("t:", t.TraceID())
		// //new c.Context
		// c.Context = trace.NewContext(c.Context, t)
		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		s, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		if nil != err {
			panic(err)
		}
		span := opentracing.StartSpan("eee", opentracing.ChildOf(s))

		c.Next()
		span.Finish()
	}
}
