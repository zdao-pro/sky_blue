package queue

import (
	"time"
)

// TaskExcutor the real excutor of the task
// if you want write a task excutor, please implement it
type TaskExcutor interface {
	Start()                             // start the excutor
	Stop()                              // stop the excutor
	SetTaskBlockedQueue(q BlockedQueue) // set the task blocked queue
	Close()                             // release the resource
}

// TaskQueue the inetrface for the task queue
type TaskQueue interface {
	SubmitTask(Task) bool
	SetParallelTaskNum(int)
	Start()
	TaskSize() int
	Close()
	Wait(time.Duration) error
}
