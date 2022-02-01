package sync

import "sync"

var waitGroupPool = sync.Pool{
	New: func() interface{} {
		return new(sync.WaitGroup)
	},
}

func AcquireWaitGroup() *sync.WaitGroup {
	return waitGroupPool.Get().(*sync.WaitGroup) // nolint:forcetypeassert
}

func ReleaseWaitGroup(wg *sync.WaitGroup) {
	waitGroupPool.Put(wg)
}
