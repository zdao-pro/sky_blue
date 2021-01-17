package trace

import (

	// "github.com/zdao-pro/sky_blue/pkg/log"

	"github.com/opentracing/opentracing-go"
)

var (
	_tracer opentracing.Tracer
)

//SetGlobalTracer ..
func SetGlobalTracer(tracer opentracing.Tracer) {
	_tracer = tracer
	opentracing.SetGlobalTracer(_tracer)
}
