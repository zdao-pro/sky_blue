package gin

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

//LogConfig log configure
type LogConfig struct {
	OutPut io.Writer
}

//ParamData ...
type ParamData struct {
	Request  *http.Request
	TimeNow  time.Time
	Latency  time.Duration
	HTTPCode int
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
}

func accessRender(param ParamData) string {
	return fmt.Sprintf("%v %v %d %s %s %s %s", param.TimeNow, param.Latency.Milliseconds(), param.HTTPCode, param.ClientIP, param.Method, param.Path, param.ErrorMessage)
}

//GetAccessLogger retunrn access logger
func GetAccessLogger(conf LogConfig) HandlerFunc {
	out := conf.OutPut
	if nil != out {
		DefaultWriter = out
	}

	return func(c *Context) {
		param := ParamData{
			Request: c.Request,
		}

		start := time.Now()

		c.Next()

		param.TimeNow = time.Now()
		param.Latency = param.TimeNow.Sub(start)
		param.Path = c.Request.URL.Path
		param.Method = c.Request.Method
		param.HTTPCode = c.Writer.Status()
		param.ClientIP = c.ClientIP()
		log.Access(accessRender(param))
	}
}
