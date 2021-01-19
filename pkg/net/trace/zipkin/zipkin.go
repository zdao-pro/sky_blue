package zipkin

import (
	"os"

	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"

	"github.com/openzipkin/zipkin-go/idgenerator"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"

	// zipkinkafka "github.com/openzipkin/zipkin-go/reporter/kafka"
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

// func tracerOptions(opts *zkOt.TracerOptions) {
// 	opts. = zkOt.B3InjectSingle
// }

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
		// r, err := zipkinkafka.NewReporter([]string{config.KafkaHost})
		// if nil != err {
		// 	panic(err)
		// }
		// Reporter = r
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

	t := zkOt.Wrap(nativeTracer) // zkOt.WithB3InjectOption(zkOt.B3InjectStandard)
	trace.SetGlobalTracer(t)
}

//GenerateTraceID ..
func GenerateTraceID() string {
	i := idgenerator.NewRandom64()
	tr := i.TraceID()
	return tr.String()
}

//Close ..
func Close() {
	Reporter.Close()
}
