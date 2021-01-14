package log

import (
	"context"
	"fmt"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/net/trace"
)

//Handles ...
type Handles struct {
	handles []Handle
}

func newHandles(hans ...Handle) Handle {
	th := Handles{
		handles: hans,
	}
	return th
}

//Log ...
func (hs Handles) Log(ctx context.Context, l Level, d ...D) {
	//assample the D slice
	ds := make([]D, 0, 5)
	if _accessLevel == l {
		ds = append(ds, d...)
	} else {
		ds = append(ds, prefixD[l], KV(_time, time.Now()), KVString(_level, l.String()))
		ds = append(ds, d...)
		if logConfig.Source {
			fn := funcName(3)
			ds = append(ds, KVString(_source, fn))
		}
	}
	if t, ok := trace.FromContext(ctx); ok {
		ds = append(ds, KVString(_log, fmt.Sprintf("\x1b[97;41mtrace_id:%s\x1b[0m", t.TraceID())))
	}
	ds = append(ds, tailD[l])

	for _, f := range hs.handles {
		f.Log(ctx, l, ds...)
	}
}

//Close ...
func (hs Handles) Close() error {
	for _, f := range hs.handles {
		if err := f.Close; nil != err {

		}
	}
	return nil
}
