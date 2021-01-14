package trace

//GinTrace ..
type GinTrace struct {
	traceID string
}

//TraceID ..
func (g *GinTrace) TraceID() string {
	return g.traceID
}

//NewGinTrace ..
func NewGinTrace(s string) Trace {
	return &GinTrace{
		traceID: s,
	}
}
