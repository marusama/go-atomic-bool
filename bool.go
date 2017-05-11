package atomic

import "sync/atomic"

// AtomicBool atomic bool
type AtomicBool struct {
	x int32
}

// NewAtomicBool new AtomicBool
func NewAtomicBool(b bool) *AtomicBool {
	var i int32
	if b {
		i = 1
	}
	return &AtomicBool{
		x: i,
	}
}

// Get get bool value
func (this *AtomicBool) Get() bool {
	return atomic.LoadInt32(&(this.x)) != 0
}

// Set set bool value
func (this *AtomicBool) Set(b bool) {
	var i int32
	if b {
		i = 1
	}
	atomic.StoreInt32(&(this.x), int32(i))
}
