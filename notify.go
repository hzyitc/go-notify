package notify

import "sync/atomic"

type Notify struct {
	c atomic.Value
}

// Return a chan which will be closed when calling `n.NotifyAll()`
func (n *Notify) Wait() <-chan interface{} {
	c := n.c.Load()
	if c == nil {
		n.c.CompareAndSwap(nil, make(chan interface{}))
		c = n.c.Load()
	}
	return c.(chan interface{})
}

// Notify all waiting goroutines
func (n *Notify) NotifyAll() {
	old := n.c.Swap(make(chan interface{}))
	if old != nil {
		close(old.(chan interface{}))
	}
}
