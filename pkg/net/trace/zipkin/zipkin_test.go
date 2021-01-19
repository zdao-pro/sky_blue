package zipkin

import (
	"net/http"
	"testing"
	"time"

	// zipkinkafka "github.com/openzipkin/zipkin-go/reporter/kafka"
	"github.com/opentracing/opentracing-go"
	// "github.com/Shopify/sarama"
)

func TestZipkin(t *testing.T) {
	// r, err := zipkinkafka.NewReporter([]string{"10.20.2.156:9092"})
	// if nil != err {
	// 	panic(err)
	// }
	// defer r.Close()

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
