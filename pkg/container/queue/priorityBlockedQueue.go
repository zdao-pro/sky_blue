package queue

import (
	"container/heap"
	"context"
	"time"
)

// PriorityBlockedQueue ..
type PriorityBlockedQueue struct {
	queue    *ArrayQueue
	notEmpty chan struct{}
}

// NewPriorityBlockedQueue ..
func NewPriorityBlockedQueue() *PriorityBlockedQueue {
	return &PriorityBlockedQueue{
		queue:    &ArrayQueue{},
		notEmpty: make(chan struct{}, 300),
	}
}

// Push push the element
func (b *PriorityBlockedQueue) Push(ctx context.Context, v interface{}) error {
	heap.Push(b.queue, v)
	b.notEmpty <- struct{}{}
	return nil
}

// Pop pop the element
func (b *PriorityBlockedQueue) Pop(timeOut time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	select {
	case <-b.notEmpty:
	case <-ctx.Done():
		return nil, ErrWaitTimeOut
	}

	if e := heap.Pop(b.queue); e != nil {
		return e, nil
	}
	return nil, ErrWaitTimeOut
}

// Len return the length of the queue
func (b *PriorityBlockedQueue) Len() int {
	return b.queue.Len()
}
