package queue

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"

	"encoding/gob"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertQueue(t *testing.T) {
	Convey("len", t, func() {
		q := NewBlockedLinkQueue()
		So(q.Len(), ShouldEqual, 0)
		ctx := context.Background()
		q.Push(ctx, 10)
		q.Push(ctx, 100)
		So(q.Len(), ShouldEqual, 2)
		i, err := q.Pop(time.Millisecond * 5)
		So(err, ShouldEqual, nil)
		So(i, ShouldEqual, 10)

		b, err := q.Pop(time.Millisecond * 5)
		So(err, ShouldEqual, nil)
		So(b, ShouldEqual, 100)

		c, err := q.Pop(time.Millisecond * 5)
		So(c, ShouldBeNil)
		So(err, ShouldNotBeNil)

		d, err := q.Pop(time.Second * 20)
		So(d, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestInsertQueueConcurrent(t *testing.T) {
	Convey("insertConcurrently", t, func() {
		q := NewBlockedLinkQueue()
		ctx := context.Background()
		var num int64 = 3000
		sema := semaphore.NewWeighted(num)
		sema.Acquire(ctx, num)
		for i := 1; i <= int(num); i++ {
			go func(j int) {
				// fmt.Println("j:", j)
				q.Push(ctx, j)
				sema.Release(1)
			}(i)
		}
		sema.Acquire(ctx, num)
		So(q.Len(), ShouldEqual, num)
		var sum int
		for i := 1; i <= int(num); i++ {
			b, err := q.Pop(time.Millisecond * 5)
			So(err, ShouldBeNil)
			// fmt.Println("b:", b)
			sum = sum + b.(int)
		}
		in := int(num)
		So(sum, ShouldEqual, (1+in)*in/2)
	})
}

func TestTaskQueueInit(t *testing.T) {
	Convey("init TaskQueue", t, func() {
		num := 10
		q := NewDefaultTaskQueue(num)
		So(q, ShouldNotBeNil)
	})
}

type TaskMock struct {
	I int
}

func (m *TaskMock) Run() error {
	time.Sleep(time.Second)
	fmt.Println("hhhhhhhhhhhhhhhh:", m.I)
	return nil
}

func NewTaskMock(a int) *TaskMock {
	return &TaskMock{
		I: a,
	}
}

func TestTaskQueueInsert(t *testing.T) {
	Convey("init TaskQueue Insert", t, func() {
		num := 1
		q := NewDefaultTaskQueue(num)
		So(q, ShouldNotBeNil)
		t1 := NewTaskMock(1)
		q.SubmitTask(t1)

		t2 := NewTaskMock(2)
		q.SubmitTask(t2)

		t3 := NewTaskMock(3)
		q.SubmitTask(t3)

		t4 := NewTaskMock(4)
		q.SubmitTask(t4)

		q.Start()
		err := q.Wait(5 * time.Second)
		So(err, ShouldNotBeNil)
	})
}

func TestTaskQueueStop(t *testing.T) {
	Convey("init TaskQueue Insert", t, func() {
		num := 1
		q := NewDefaultTaskQueue(num)
		So(q, ShouldNotBeNil)
		t1 := NewTaskMock(1)
		q.SubmitTask(t1)

		t2 := NewTaskMock(2)
		q.SubmitTask(t2)
		q.Start()
		time.Sleep(2 * time.Second)
		q.Close()
		t3 := NewTaskMock(3)
		q.SubmitTask(t3)

		t4 := NewTaskMock(4)
		q.SubmitTask(t4)

		err := q.Wait(7 * time.Second)
		So(err, ShouldNotBeNil)

	})
}

func TestGob(t *testing.T) {
	Convey("gob", t, func() {
		var t1 TaskMock
		t1.I = 90
		var w bytes.Buffer
		enc := gob.NewEncoder(&w)
		err := enc.Encode(&t1)
		So(err, ShouldBeNil)

		dec := gob.NewDecoder(&w)
		var t2 TaskMock
		err = dec.Decode(&t2)
		So(err, ShouldBeNil)
		if err == nil {
			t2.Run()
		}
	})
}

type PriorityTaskMock struct {
	Score int
}

func (m *PriorityTaskMock) Run() error {
	time.Sleep(time.Second)
	fmt.Println("hhhhhhhhhhhhhhhh,score:", m.Score)
	return nil
}

// Less reports whether the element with
// index i should sort before self.
func (m *PriorityTaskMock) Less(i interface{}) bool {
	if p, ok := i.(*PriorityTaskMock); ok {
		return m.Score < p.Score
	}
	return false
}

func NewPriorityTaskMock(a int) *PriorityTaskMock {
	return &PriorityTaskMock{
		Score: a,
	}
}

func TestPriorityTaskQueue(t *testing.T) {
	Convey("TestPriorityTaskQueue", t, func() {
		q := NewPriorityBlockedQueue()
		So(q, ShouldNotBeNil)
		t1 := NewPriorityTaskMock(12)

		ctx := context.Background()
		q.Push(ctx, t1)

		So(q.Len(), ShouldEqual, 1)

		_, err := q.Pop(5 * time.Millisecond)
		So(q.Len(), ShouldEqual, 0)
		So(err, ShouldBeNil)
	})
}

func TestPriorityTaskQueueInsert(t *testing.T) {
	Convey("init TaskQueue Insert", t, func() {
		num := 1
		q := NewDefaultPriorityTaskQueue(num)
		So(q, ShouldNotBeNil)
		t1 := NewPriorityTaskMock(67)
		q.SubmitPriorityTask(t1)

		t2 := NewPriorityTaskMock(2)
		q.SubmitPriorityTask(t2)

		t3 := NewPriorityTaskMock(3)
		q.SubmitPriorityTask(t3)

		t4 := NewPriorityTaskMock(4)
		q.SubmitPriorityTask(t4)

		q.Start()
		err := q.Wait(5 * time.Second)
		So(err, ShouldNotBeNil)
	})
}
