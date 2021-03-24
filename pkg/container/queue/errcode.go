package queue

import (
	"errors"
)

var (
	// ErrNoElement the queue has no element
	ErrNoElement = errors.New("the queue has no element")
	// ErrWaitTimeOut ..
	ErrWaitTimeOut = errors.New("wait time out")
	// ErrNotComparator the element not implement the Comparator
	ErrNotComparator = errors.New("the element not implement the Comparator")
)
