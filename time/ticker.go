//go:build go1.15
// +build go1.15

package time

import (
	"sync"
	"time"
)

var tickerPool = sync.Pool{}

// AcquireTicker returns a ticker from the pool if possible.
func AcquireTicker(d time.Duration) *time.Ticker {
	v := tickerPool.Get()
	if v == nil {
		return time.NewTicker(d)
	}

	t, ok := v.(*time.Ticker)
	if !ok {
		panic("unexpected type of time.Ticker")
	}

	t.Reset(d)

	return t
}

// ReleaseTicker returns a ticker into the pool.
func ReleaseTicker(tm *time.Ticker) {
	tm.Stop()
	tickerPool.Put(tm)
}
