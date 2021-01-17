package trace

import (
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
)

func TestZipkin(t *testing.T) {
	Init()
	trace := NewTracer("Main")
	span := trace.StartSpan("zipkin_test")
	time.Sleep(time.Duration(1) * time.Second)
	spa2 := opentracing.StartSpan("zip", opentracing.ChildOf(span.Context()))
	spa3 := opentracing.StartSpan("zip2", opentracing.ChildOf(span.Context()))

	time.Sleep(time.Duration(1) * time.Second)
	spa3.Finish()
	time.Sleep(time.Duration(2) * time.Second)
	spa2.Finish()
	time.Sleep(time.Duration(1) * time.Second)
	span.Finish()

}
