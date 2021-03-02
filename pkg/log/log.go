package log

import (
	"context"
	"fmt"
	"os"

	"github.com/zdao-pro/sky_blue/pkg/env"
)

var (
	h Handle
)

//Config ...
type Config struct {
	_debugPrintFlag bool
	_infoPrintFlag  bool
	_warnPrintFlag  bool
	_errorPrintFlag bool
	_fetalPrintFlag bool
	// Filter tell log handler which field are sensitive message, use * instead.
	Filter []string
	// file
	Dir string
	// stdout
	Stdout bool
	Nlog   bool
	Source bool
}

var logConfig Config

func init() {
	Init(nil)
}

//Init ...
func Init(conf *Config) {
	if nil == conf {
		logConfig = Config{
			_infoPrintFlag:  true,
			_warnPrintFlag:  true,
			_errorPrintFlag: true,
			_fetalPrintFlag: true,
			_debugPrintFlag: true,
			Stdout:          true,
			Source:          true,
		}
	} else {
		logConfig = *conf
	}
	logConfig._infoPrintFlag = true
	logConfig._debugPrintFlag = true
	logConfig._errorPrintFlag = true
	logConfig._fetalPrintFlag = true
	logConfig._warnPrintFlag = true
	if env.IsOnline() {
		logConfig._debugPrintFlag = false
	}

	// if !env.IsDev() {
	// 	logConfig.Stdout = true
	// }
	var hs []Handle

	//udp log handle
	udpServerAddr := os.Getenv("UDP_LOG_ADDR")
	udpServerPort := os.Getenv("UDP_LOG_PORT")
	if "" != udpServerAddr && "" != udpServerPort {
		nlogH := newNlogHnadle(udpServerAddr, udpServerPort)
		hs = append(hs, nlogH)
	}

	if logConfig.Stdout {
		hs = append(hs, newStdoutHandle())
	}

	h = newHandles(hs...)
}

//Debug ...
func Debug(format string, args ...interface{}) {
	if logConfig._debugPrintFlag {
		h.Log(context.Background(), _debugLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Info ...
func Info(format string, args ...interface{}) {
	if logConfig._infoPrintFlag {
		h.Log(context.Background(), _infoLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Warn ...
func Warn(format string, args ...interface{}) {
	if logConfig._warnPrintFlag {
		h.Log(context.Background(), _warnLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Error ...
func Error(format string, args ...interface{}) {
	if logConfig._errorPrintFlag {
		h.Log(context.Background(), _errorLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Fetal ...
func Fetal(format string, args ...interface{}) {
	if logConfig._fetalPrintFlag {
		h.Log(context.Background(), _fetalLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Access ...
func Access(format string, args ...interface{}) {
	if logConfig._fetalPrintFlag {
		h.Log(context.Background(), _accessLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Debugc ...
func Debugc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._debugPrintFlag {
		h.Log(ctx, _debugLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Infoc ...
func Infoc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._infoPrintFlag {
		h.Log(ctx, _infoLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Warnc ...
func Warnc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._warnPrintFlag {
		h.Log(ctx, _warnLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Error ...
func Errorc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._errorPrintFlag {
		h.Log(ctx, _errorLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Fetal ...
func Fetalc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._fetalPrintFlag {
		h.Log(ctx, _fetalLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}

//Access ...
func Accessc(ctx context.Context, format string, args ...interface{}) {
	if logConfig._fetalPrintFlag {
		h.Log(ctx, _accessLevel, KVString(_log, fmt.Sprintf(format, args...)))
	}
}
