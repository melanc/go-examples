package queue

import (
	"sync"

	"gopkg.in/eapache/queue.v1"
)

// Synchronous FIFO queue
type SyncQueue struct {
	lock   sync.Mutex
	cond   *sync.Cond
	buffer *queue.Queue
}

// Create a new SyncQueue
func NewSyncQueue() *SyncQueue {
	q := &SyncQueue{
		buffer: queue.New(),
	}
	q.cond = sync.NewCond(&q.lock)
	return q
}

// Pop an item from SyncQueue, will block if SyncQueue is empty
func (q *SyncQueue) Pop() interface{} {
	c := q.cond
	buffer := q.buffer

	q.lock.Lock()
	for buffer.Length() == 0 {
		c.Wait()
	}

	v := buffer.Peek()
	buffer.Remove()

	q.lock.Unlock()
	return v
}

// Push an item to SyncQueue. Always returns immediately without blocking
func (q *SyncQueue) Push(v interface{}) {
	q.lock.Lock()

	q.buffer.Add(v)
	q.cond.Signal()

	q.lock.Unlock()
}

// Get the length of SyncQueue
func (q *SyncQueue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.buffer.Length()
}
