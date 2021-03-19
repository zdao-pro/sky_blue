package queue

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"

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
	i int
}

func (m *TaskMock) Run() error {
	time.Sleep(time.Second)
	fmt.Println("hhhhhhhhhhhhhhhh:", m.i)
	return nil
}

func NewTaskMock(a int) *TaskMock {
	return &TaskMock{
		i: a,
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
		time.Sleep(2*time.Second)
		q.Close()
		t3 := NewTaskMock(3)
		q.SubmitTask(t3)

		t4 := NewTaskMock(4)
		q.SubmitTask(t4)

		err := q.Wait(7 * time.Second)
		So(err, ShouldNotBeNil)
		
	})
}
