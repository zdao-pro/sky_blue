package zipkin

import (
	"os"

	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/zdao-pro/sky_blue/pkg/net/trace"
	"github.com/zdao-pro/sky_blue/pkg/util"
)

// 第一步: 开一个全局变量
var zkTracer opentracing.Tracer

//Reporter ..
var Reporter reporter.Reporter

//Config is...
type Config struct {
	ZipkinHost string
}

var config Config

// Init ..
func Init(serviveName string) {
	config = Config{
		ZipkinHost: "http://zipkin.zhaodao88.com/api/v2/spans",
	}
	zipkinHost := os.Getenv("ZIPKIN_HOS")
	if "" != zipkinHost {
		config.ZipkinHost = zipkinHost
	}
	// fmt.Println(util.GetLocalAddress())
	Reporter = zkHttp.NewReporter(config.ZipkinHost)
	endpoint, err := zipkin.NewEndpoint(serviveName, util.GetLocalAddress())
	if err != nil {
		panic(err)
	}

	nativeTracer, err := zipkin.NewTracer(Reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}

	t := zkOt.Wrap(nativeTracer)
	trace.SetGlobalTracer(t)
}

//Close ..
func Close() {
	Reporter.Close()
}
