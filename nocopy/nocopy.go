package nocopy

// Lock is a no-op implementation of the sync.Locker interface.
func (*NoCopy) Lock() {}

// Unlock is a no-op implementation of the sync.Locker interface.
func (*NoCopy) Unlock() {}
