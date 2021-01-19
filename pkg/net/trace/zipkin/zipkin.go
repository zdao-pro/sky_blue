package zipkin

import (
	"os"

	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"

	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	zipkinkafka "github.com/openzipkin/zipkin-go/reporter/kafka"
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
	KafkaHost  string
}

var config Config

// Init ..
func Init(serviveName string) {
	zipkinHost := os.Getenv("ZIPKIN_HOST")
	if "" != zipkinHost {
		config.ZipkinHost = zipkinHost
	}

	//@example: '127.0.0.1:9092'
	kafkaHost := os.Getenv("KAFKA_HOST")
	if "" != kafkaHost {
		config.KafkaHost = kafkaHost
	}

	// fmt.Println(util.GetLocalAddress())

	if config.ZipkinHost != "" {
		Reporter = zkHttp.NewReporter(config.ZipkinHost)
	} else if config.KafkaHost != "" {
		r, err := zipkinkafka.NewReporter([]string{config.KafkaHost})
		if nil != err {
			panic(err)
		}
		Reporter = r
	} else {
		return
	}

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
