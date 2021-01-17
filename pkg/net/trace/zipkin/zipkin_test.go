package zipkin

import (
	"net/http"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
)

func TestZipkin(t *testing.T) {
	{
		var r http.Request
		Init("hhaha")
		defer Close()
		carrier := opentracing.HTTPHeadersCarrier(r.Header)
		s, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		if nil != err {
			panic(err)
		}
		span := opentracing.StartSpan("eee", opentracing.ChildOf(s))
		time.Sleep(time.Duration(1) * time.Second)
		span.Finish()
		// reporter := zkHttp.NewReporter("http://zipkin.zhaodao88.com/api/v2/spans")
		// defer reporter.Close()
		// endpoint, err := zipkin.NewEndpoint("main3", "localhost")
		// if err != nil {
		// 	log.Fatalf("unable to create local endpoint: %+v\n", err)
		// }
		// nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		// if err != nil {
		// 	log.Fatalf("unable to create tracer: %+v\n", err)
		// }
		// zkTracer = zkOt.Wrap(nativeTracer)
		// opentracing.SetGlobalTracer(zkTracer)

		// span := zkTracer.StartSpan("zipkin_test")
		// time.Sleep(time.Duration(1) * time.Second)
		// spa2 := opentracing.StartSpan("zip", opentracing.ChildOf(span.Context()))
		// spa3 := opentracing.StartSpan("zip2", opentracing.FollowsFrom(spa2.Context()))
		// time.Sleep(time.Duration(1) * time.Second)
		// spa3.Finish()
		// time.Sleep(time.Duration(2) * time.Second)
		// spa2.Finish()
		// time.Sleep(time.Duration(1) * time.Second)
		// span.Finish()
	}
}
