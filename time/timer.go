//go:build go1.15
// +build go1.15

package time

import (
	"sync"
	"time"
)

var timerPool = sync.Pool{}

// AcquireTimer returns a timer from the pool if possible.
func AcquireTimer(d time.Duration) *time.Timer {
	v := timerPool.Get()
	if v == nil {
		return time.NewTimer(d)
	}

	t, ok := v.(*time.Timer)
	if !ok {
		panic("unexpected type of time.Timer")
	}

	if t.Reset(d) {
		// active timer?
		return time.NewTimer(d)
	}

	return t
}

// ReleaseTimer returns a timer into the pool.
func ReleaseTimer(t *time.Timer) {
	if !t.Stop() {
		// tm.Stop() returns false if the timer has already expired or been stopped.
		// We can't be sure that timer.C will not be filled after timer.Stop(),
		// see https://groups.google.com/forum/#!topic/golang-nuts/-8O3AknKpwk
		//
		// The tip from manual to read from timer.C possibly blocks caller if caller
		// has already done <-timer.C. Non-blocking read from timer.C with select does
		// not help either because send is done concurrently from another goroutine.
		return
	}

	timerPool.Put(t)
}
