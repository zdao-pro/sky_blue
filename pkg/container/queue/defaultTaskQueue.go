package queue

import (
	"context"
)

// DefaultTaskQueue the common implement for the task queue
type DefaultTaskQueue struct {
	taskQueue        BlockedQueue
	taskExcutorSlice []TaskExcutor
}

// NewDefaultTaskQueue ..
func NewDefaultTaskQueue(taskNum int) *DefaultTaskQueue {
	q := &DefaultTaskQueue{}
	q.SetParallelTaskNum(taskNum)
	return q
}

// SubmitTask ..
func (t *DefaultTaskQueue) SubmitTask(task Task) bool {
	if err := t.taskQueue.Push(context.Background(), task); err == nil {
		return true
	}
	return false
}

// SetParallelTaskNum ..
func (t *DefaultTaskQueue) SetParallelTaskNum(taskNum int) {
	t.taskExcutorSlice = make([]TaskExcutor, taskNum)
	for i := 1; i <= taskNum; i++ {
		t.taskExcutorSlice = append(t.taskExcutorSlice, &DefaultTaskExcutor{})
	}
}

// Start ..
func (t *DefaultTaskQueue) Start() {
	for _, excutor := range t.taskExcutorSlice {
		excutor.Start()
	}
}

// TaskSize ..
func (t *DefaultTaskQueue) TaskSize() int {
	return t.taskQueue.Len()
}

// Close ..
func (t *DefaultTaskQueue) Close() {
	for _, excutor := range t.taskExcutorSlice {
		excutor.Close()
	}
}

// Wait ..
func (t *DefaultTaskQueue) Wait(ctx context.Context) error {

}
