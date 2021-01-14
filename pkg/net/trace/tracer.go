package trace

//Trace return current trace id.
type Trace interface {
	// return current trace id.
	TraceID() string
}
