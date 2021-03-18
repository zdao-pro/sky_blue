package queue

import (
	"time"
)

// DefaultTaskExcutor ..
type DefaultTaskExcutor struct {
	taskQueue BlockedQueue
	isStop    bool
}

// NewDefaultTaskExcutor ..
func NewDefaultTaskExcutor() *DefaultTaskExcutor {
	return &DefaultTaskExcutor{}
}

// Start start the excutor
func (e *DefaultTaskExcutor) Start() {
	e.isStop = true
	go func() {
		if true == e.isStop {
			if v, err := e.taskQueue.Pop(time.Millisecond * 5); err == nil {
				task := v.(Task)
				task.Run()
			}
		}
	}()
}

// Stop stop the excutor
func (e *DefaultTaskExcutor) Stop() {
	e.isStop = false
}

// SetTaskBlockedQueue set the task blocked queue
func (e *DefaultTaskExcutor) SetTaskBlockedQueue(q BlockedQueue) {
	e.taskQueue = q
}

// Close ..
func (e *DefaultTaskExcutor) Close() {

}
