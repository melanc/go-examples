package queue

import (
	"container/list"
	"sync"
)

type FIFO struct {
	list *list.List
	lock *sync.Mutex
	cond *sync.Cond
}

// NewFIFO create a FIFO instance
func NewFIFO() *FIFO {
	list := list.New()
	lock := new(sync.Mutex)
	cond := sync.NewCond(lock)

	return &FIFO{list, lock, cond}
}

// Dequeue returns the first element of the queue
func (q *FIFO) Dequeue() interface{} {
	q.lock.Lock()
	// block when queue empty
	for q.list.Len() == 0 {
		q.cond.Wait()
	}

	// get front element of list
	e := q.list.Front()
	// remove front element
	q.list.Remove(e)

	q.lock.Unlock()
	return e.Value
}

// Enqueue inserts the element in the end of the queue
func (q *FIFO) Enqueue(v interface{}) {
	q.lock.Lock()
	// put element in the end of list
	q.list.PushBack(v)
	// notify dequeue to unblock
	q.cond.Signal()
	q.lock.Unlock()
}

// Size return the size of queue
func (q *FIFO) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.Len()
}
