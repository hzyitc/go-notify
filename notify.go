package notify

import (
	"sync/atomic"
	"unsafe"
)

type Notify struct {
	v unsafe.Pointer
}

// Return a chan which will be closed when calling `n.NotifyAll()`
func (n *Notify) Wait() <-chan interface{} {
	v := atomic.LoadPointer(&n.v)
	if v == nil {
		ch := make(chan interface{})
		atomic.CompareAndSwapPointer(&n.v, nil, unsafe.Pointer(&ch))
		v = atomic.LoadPointer(&n.v)
	}
	return *(*chan interface{})(v)
}

// Notify all waiting goroutines
func (n *Notify) NotifyAll() {
	ch := make(chan interface{})
	old := atomic.SwapPointer(&n.v, unsafe.Pointer(&ch))
	if old != nil {
		close(*(*chan interface{})(old))
	}
}
