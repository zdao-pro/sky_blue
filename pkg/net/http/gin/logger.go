package gin

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/trace"
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
	Request        *http.Request
	TimeNow        time.Time
	Latency        time.Duration
	HTTPCode       int
	HTTPCodeFormat string
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage        string
	ContentLength       int64
	Connection          int
	ConnectionRequests  int64
	RequestStr          string
	UserID              int
	ResponsonBodyLength int
	HTTPXIsErrorCode    string
	HTTPXIsErrorMsg     string
	HTTPUserAgent       string
	HTTPReferer         string
	HTTPXForwardedFor   string
	HTTPXIsIP           string
	HTTPHost            string
	RequestData         string
	HTTPXIsClientIP     string
	TraceID             string
	RequestTimeFormat   string
	RequestTime         float32
	HTTPVersion         string
}

func accessRender(param ParamData) string {
	return fmt.Sprintf("%s %s %v %s%f\x1b[0m %d %d %d %s \"%s %s %s\" \"%d\" %s%d\x1b[0m %d \"%s\" \"%s\" \"%s\" \"%s\" \"%s\" \"%s\" %s \"%s\" \"%s\" \"%s\"",
		param.ClientIP,
		param.TimeNow.Format(time.RFC3339),
		param.TimeNow.Unix(),
		param.RequestTimeFormat,
		param.RequestTime,
		param.ContentLength,
		param.Connection,
		param.ConnectionRequests,
		param.Path,
		param.Method,
		param.RequestStr,
		param.HTTPVersion,
		param.UserID,
		param.HTTPCodeFormat,
		param.HTTPCode,
		param.ResponsonBodyLength,
		param.HTTPXIsErrorCode,
		param.HTTPXIsErrorMsg,
		param.HTTPReferer,
		param.HTTPUserAgent,
		param.HTTPXForwardedFor,
		param.HTTPXIsIP,
		param.HTTPHost,
		param.RequestData,
		param.HTTPXIsClientIP,
		param.TraceID)
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
		param.HTTPXIsErrorCode = c.Writer.Header().Get("X-IS-Error-Code")
		param.HTTPXIsErrorMsg = c.Writer.Header().Get("X-IS-Error-Msg")
		param.HTTPReferer = c.GetHeader("Referer")
		param.HTTPUserAgent = c.GetHeader("User-Agent")
		param.HTTPXForwardedFor = c.GetHeader("X-Forwarded-For")
		param.HTTPXIsIP = c.GetHeader("X-IS-IP")
		param.HTTPHost = c.GetHeader("Host")
		param.ResponsonBodyLength = c.Writer.Size()
		len := param.ContentLength
		if 200 < param.ContentLength {
			param.ContentLength = 200
		}
		body := make([]byte, len)
		c.Request.Body.Read(body)
		param.RequestData = string(body)
		param.RequestData = strings.Replace(param.RequestData, "\n", "", -1)
		//decorate request time format
		param.RequestTime = float32(param.Latency.Milliseconds()) / 1000
		if param.RequestTime < 0.01 {
			param.RequestTimeFormat = "\x1b[0;00m"
		} else if param.RequestTime >= 0.01 && param.RequestTime < 0.1 {
			param.RequestTimeFormat = "\x1b[0;36;40m"
		} else if param.RequestTime >= 0.1 && param.RequestTime < 0.5 {
			param.RequestTimeFormat = "\x1b[0;33;40m"
		} else {
			param.RequestTimeFormat = "\x1b[0;31;400m"
		}

		if param.HTTPCode >= 500 {
			param.HTTPCodeFormat = "\x1b[0;31;40m"
		} else if param.HTTPCode >= 400 {
			param.HTTPCodeFormat = "\x1b[0;33;40m"
		} else if param.HTTPCode < 500 {
			param.HTTPCodeFormat = "\x1b[0;36;40m"
		} else {
			param.HTTPCodeFormat = "\x1b[0;00m"
		}
		param.HTTPVersion = c.Request.Proto
		// fmt.Println("vveev", c.Request.Header)
		// fmt.Println("vvv", c.Request.Header.Get("X-B3-Traceid"))

		// param.TraceID = c.Request.Header.Get("X-B3-Traceid")
		if t, ok := trace.FromContext(c.Context); ok {
			param.TraceID = t.TraceID()
		}
		// c.Request.Body.
		log.Access(accessRender(param))
	}
}
