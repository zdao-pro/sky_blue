package queue

import (
	"context"
	"time"
)

// BlockedQueue ..
type BlockedQueue interface {
	Push(ctx context.Context, v interface{}) error
	Pop(waitTime time.Duration) (interface{}, error)
	Len() int
}
