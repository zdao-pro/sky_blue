package queue

import (
	"fmt"
)

var (
	// ErrNoElement the queue has no element
	ErrNoElement = fmt.Errorf("the queue has no element")
	// ErrWaitTimeOut ..
	ErrWaitTimeOut = fmt.Errorf("wait time out")
)
