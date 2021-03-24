package queue

import (
	"sync"
	"sync/atomic"
)

// ArrayQueue ..
type ArrayQueue struct {
	object []interface{}
	size   int32
	mu     sync.Mutex
}

// Push .. add x as element Len()
func (p *ArrayQueue) Push(x interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.object = append(p.object, x)
	atomic.AddInt32(&p.size, 1)
}

// Pop remove and return element Len() - 1.
func (p *ArrayQueue) Pop() interface{} {
	size := atomic.LoadInt32(&p.size)
	if size <= 0 {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	// get the last element
	e := p.object[size-1]
	// release the memorary
	p.object[size-1] = nil
	atomic.AddInt32(&p.size, -1)
	return e
}

// Len is the number of elements in the collection.
func (p *ArrayQueue) Len() int {
	return int(atomic.LoadInt32(&p.size))
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p *ArrayQueue) Less(i, j int) bool {
	ci, ok := p.object[i].(Comparator)
	if !ok {
		panic(ErrNotComparator)
	}

	return ci.Less(p.object[j])
}

// Swap swaps the elements with indexes i and j.
func (p *ArrayQueue) Swap(i, j int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.object[i], p.object[j] = p.object[j], p.object[i]
}
