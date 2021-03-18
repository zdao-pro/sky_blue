package queue

// @author: jim_sun
// @function: the block queue

import (
	"container/list"
	"context"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// BlockedLinkQueue 链表阻塞队列
type BlockedLinkQueue struct {
	list *list.List
	mu   *sync.Mutex
	sema *semaphore.Weighted
}

// Push push the element
func (b *BlockedLinkQueue) Push(ctx context.Context, v interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.list.PushFront(v)
	b.sema.Release(1)
	return nil
}

// Pop pop the element
func (b *BlockedLinkQueue) Pop(waitTime time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	err := b.sema.Acquire(ctx, 1)
	if err == nil {
		b.mu.Lock()
		defer b.mu.Unlock()
		if e := b.pop(); err != nil {
			return e.Value, nil
		}
	}
	return nil, ErrWaitTimeOut
}

// Len return the length of the queue
func (b *BlockedLinkQueue) Len() int {
	return b.list.Len()
}

func (b *BlockedLinkQueue) pop() *list.Element {
	e := b.list.Front()
	if e != nil {
		b.list.Remove(e)
	}
	return e
}
