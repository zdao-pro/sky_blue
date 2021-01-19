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
		// fmt.Println(carrier)
		s, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		if nil != err {
			panic(err)
		}

		span := opentracing.StartSpan(c.Request.URL.Path, opentracing.ChildOf(s))
		defer span.Finish()
		span.SetTag("url", c.Request.URL.Path)
		span.SetTag("param", c.Request.URL.String())
		// fmt.Println("rr:", span.BaggageItem("TraceId"))

		er := opentracing.GlobalTracer().Inject(s, opentracing.HTTPHeaders, carrier)
		if nil == er {
			// fmt.Println("vvv", c.Request.Header)
		}
		c.Context = opentracing.ContextWithSpan(c.Context, span)
		c.Next()

	}
}
