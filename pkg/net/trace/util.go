package trace

import (
	"context"
)

type ctxKey string

var _ctxkey ctxKey = "zdao_bacend_trace"

// NewContext new a trace context.
// NOTE: This method is not thread safe.
func NewContext(ctx context.Context, t Trace) context.Context {
	return context.WithValue(ctx, _ctxkey, t)
}

// FromContext returns the trace bound to the context, if any.
func FromContext(ctx context.Context) (t Trace, ok bool) {
	t, ok = ctx.Value(_ctxkey).(Trace)
	return
}
