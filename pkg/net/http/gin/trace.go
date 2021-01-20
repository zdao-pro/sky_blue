package gin

import (
	"github.com/opentracing/opentracing-go"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/trace"
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
			log.Fetalc(c.Context, err.Error())
		}
		if s != nil {
			span := opentracing.StartSpan(c.Request.URL.Path, opentracing.ChildOf(s))
			defer span.Finish()
			span.SetTag("url", c.Request.URL.Path)
			span.SetTag("param", c.Request.URL.String())
			// fmt.Println("rr:", span.BaggageItem("TraceId"))

			er := opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
			if nil != er {
				log.Warnc(c.Context, er.Error())
				// fmt.Println("vvv", c.Request.Header, "url:", c.Request.URL.Path)
			} else {
				traceID := c.Request.Header.Get("X-B3-Traceid")
				t := trace.NewGinTrace(traceID)
				//new c.Context
				c.Context = trace.NewContext(c.Context, t)
			}
			c.Context = opentracing.ContextWithSpan(c.Context, span)
		}

		c.Next()

	}
}
