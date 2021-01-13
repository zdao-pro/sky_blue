package gin

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
)

const (
	purple = "\x1b[1;35m" //紫色
	yellow = "\x1b[1;33m"
	red    = "\x1b[97;41m"
	white  = "\x1b[0;00m"
	green  = "\x1b[1;32m"
	reset  = "\033[0m"
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
	ErrorMessage       string
	ContentLength      int64
	Connection         int
	ConnectionRequests int64
	RequestStr         string
}

func accessRender(param ParamData) string {
	return fmt.Sprintf("%s %s %v %v %d %d %d %s %s %d %s %s",
		param.ClientIP,
		param.TimeNow.Format("2006-01-02 15:04:05"),
		param.TimeNow.Unix(),
		param.Latency.Milliseconds()/1000,
		param.ContentLength,
		param.Connection,
		param.ConnectionRequests,
		param.Path,
		param.RequestStr,
		param.HTTPCode,
		param.Method,
		param.ErrorMessage)
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
		param.ContentLength = c.Request.ContentLength
		param.RequestStr = c.Request.URL.String()
		c.Writer.Header().Get("")
		log.Access(accessRender(param))
	}
}
g