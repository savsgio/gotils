package rand

import (
	"math/rand"
	"sync"
	"time"
)

var sourcePool = sync.Pool{
	New: func() interface{} {
		return rand.NewSource(time.Now().UnixNano())
	},
}

func AcquireSource() rand.Source {
	return sourcePool.Get().(rand.Source) // nolint:forcetypeassert
}

func ReleaseSource(src rand.Source) {
	sourcePool.Put(src)
}
