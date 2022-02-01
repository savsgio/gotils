package rand

import (
	"math/rand"
	"sync"
)

var sourcePool = sync.Pool{
	New: func() interface{} {
		return rand.NewSource(0)
	},
}

func AcquireSource(seed int64) rand.Source {
	src := sourcePool.Get().(rand.Source) // nolint:forcetypeassert
	src.Seed(seed)

	return src
}

func ReleaseSource(src rand.Source) {
	sourcePool.Put(src)
}
