package queue

import (
	"context"
	"sync"
	"time"
)

// DefaultTaskQueue the common implement for the task queue
type DefaultTaskQueue struct {
	taskQueue BlockedQueue
	wg        sync.WaitGroup
	isStop    bool
	taskNum   int
}

// NewDefaultTaskQueue ..
func NewDefaultTaskQueue(taskNum int) *DefaultTaskQueue {
	q := &DefaultTaskQueue{
		taskQueue: NewBlockedLinkQueue(),
	}
	q.SetParallelTaskNum(taskNum)
	return q
}

// NewDefaultPriorityTaskQueue ..
func NewDefaultPriorityTaskQueue(taskNum int) *DefaultTaskQueue {
	q := &DefaultTaskQueue{
		taskQueue: NewPriorityBlockedQueue(),
	}
	q.SetParallelTaskNum(taskNum)
	return q
}

// SubmitPriorityTask ..
func (t *DefaultTaskQueue) SubmitPriorityTask(task PriorityTask) bool {
	if err := t.taskQueue.Push(context.Background(), task); err == nil {
		t.wg.Add(1)
		return true
	}
	return false
}

// SubmitTask ..
func (t *DefaultTaskQueue) SubmitTask(task Task) bool {
	if err := t.taskQueue.Push(context.Background(), task); err == nil {
		t.wg.Add(1)
		return true
	}
	return false
}

// SetParallelTaskNum ..
func (t *DefaultTaskQueue) SetParallelTaskNum(taskNum int) {
	t.taskNum = taskNum
}

// Start ..
func (t *DefaultTaskQueue) Start() {
	t.isStop = true
	for i := 1; i <= t.taskNum; i++ {
		go func() {
			for {
				if true == t.isStop {
					if v, err := t.taskQueue.Pop(time.Millisecond * 5); err == nil {
						task := v.(Task)
						task.Run()
						t.wg.Done()
					}
				}
			}
		}()
	}
}

// TaskSize ..
func (t *DefaultTaskQueue) TaskSize() int {
	return t.taskQueue.Len()
}

// Close ..
func (t *DefaultTaskQueue) Close() {
	t.isStop = false
}

// Wait ..
func (t *DefaultTaskQueue) Wait(timeOut time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	go func() {
		t.wg.Wait()
	}()
	select {
	case <-ctx.Done():
		return ErrWaitTimeOut
	}
}
