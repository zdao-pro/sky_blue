package queue

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayQueue(t *testing.T) {
	Convey("array queue", t, func() {
		q := &ArrayQueue{}
		q.Push(1)
		So(q.Len(), ShouldEqual, 1)
		q.Push(2)
		So(q.Len(), ShouldEqual, 2)

		e := q.Pop()
		So(e, ShouldEqual, 2)
		So(q.Len(), ShouldEqual, 1)

		b := q.Pop()
		So(b, ShouldEqual, 1)
		So(q.Len(), ShouldEqual, 0)
	})
}

func TestArrayQueueConccurent(t *testing.T) {
	Convey("array concurrent queue", t, func() {
		q := &ArrayQueue{}
		num := 4000
		for i := 1; i <= num; i++ {
			go func(j int) {
				q.Push(j)
			}(i)
		}
		time.Sleep(5 * time.Second)
		So(q.Len(), ShouldEqual, num)

		var count int
		for i := 1; i <= num; i++ {
			if b, ok := q.Pop().(int); ok {
				count = count + b
			}
		}

		So(count, ShouldEqual, (1+num)*num/2)
	})
}

func TestArrayQueueSwap(t *testing.T) {
	Convey("swap", t, func() {
		q := &ArrayQueue{}
		q.Push(NewElement(30))
		q.Push(NewElement(20))

		So(q.Less(0, 1), ShouldEqual, false)
		So(q.Less(1, 0), ShouldEqual, true)
	})
}

type Element struct {
	I int
}

func NewElement(score int) *Element {
	return &Element{
		I: score,
	}
}

func (e *Element) Less(i interface{}) bool {
	if b, ok := i.(*Element); ok {
		return e.I < b.I
	}
	panic("gg")
}
