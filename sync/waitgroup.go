package sync

import "sync"

var waitGroupPool = sync.Pool{
	New: func() any {
		return new(sync.WaitGroup)
	},
}

// AcquireWaitGroup returns a WaitGroup from the pool.
func AcquireWaitGroup() *sync.WaitGroup {
	return waitGroupPool.Get().(*sync.WaitGroup) // nolint:forcetypeassert
}

// ReleaseWaitGroup returns a WaitGroup to the pool.
func ReleaseWaitGroup(wg *sync.WaitGroup) {
	waitGroupPool.Put(wg)
}
