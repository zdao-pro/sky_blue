package queue

// @author: jim_sun
// @function: the block queue

import (
	"container/list"
	"context"
	"sync"
	"time"
)

// BlockedLinkQueue 链表阻塞队列
type BlockedLinkQueue struct {
	list     *list.List
	mu       sync.Mutex
	notEmpty chan struct{}
}

// NewBlockedLinkQueue build the NewBlockedLinkQueue pointer
func NewBlockedLinkQueue() *BlockedLinkQueue {
	return &BlockedLinkQueue{
		list:     list.New(),
		notEmpty: make(chan struct{}, 300),
	}
}

// Push push the element
func (b *BlockedLinkQueue) Push(ctx context.Context, v interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.list.PushFront(v)
	b.notEmpty <- struct{}{}
	return nil
}

// Pop pop the element
func (b *BlockedLinkQueue) Pop(timeOut time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	select {
	case <-b.notEmpty:
	case <-ctx.Done():
		return nil, ErrWaitTimeOut
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if e := b.pop(); e != nil {
		return e.Value, nil
	}
	return nil, ErrWaitTimeOut
}

// Len return the length of the queue
func (b *BlockedLinkQueue) Len() int {
	return b.list.Len()
}

func (b *BlockedLinkQueue) pop() *list.Element {
	e := b.list.Back()
	if e != nil {
		b.list.Remove(e)
	}
	return e
}
