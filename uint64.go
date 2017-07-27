package atoms

import "sync/atomic"

// Uint64 is an atomic uint64
type Uint64 struct {
	v uint64
}

// Load will get the current value
func (u *Uint64) Load() (n uint64) {
	return atomic.LoadUint64(&u.v)
}

// Add will increment the current value by n
func (u *Uint64) Add(n uint64) (new uint64) {
	return atomic.AddUint64(&u.v, n)
}

// Store will perform an atomic store for a new value
func (u *Uint64) Store(new uint64) {
	atomic.StoreUint64(&u.v, new)
}

// Swap will perform an atomic swap for a new value
func (u *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(&u.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint64) CompareAndSwap(old, new uint64) (changed bool) {
	return atomic.CompareAndSwapUint64(&u.v, old, new)
}