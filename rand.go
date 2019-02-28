// Package rand implements fast pesudorandom number generator
// that should scale well on multi-CPU systems.
//
// Use crypto/rand instead of this package for generating
// cryptographically secure random numbers.
package rand

import (
	"sync/atomic"
	"time"
)

var rndn = getRandomInt32()

// Uint32 returns a tread-safe, non-cryptographic pseudorandom uint32.
func Uint32() uint32 {
	x := uint32(atomic.AddInt32(&rndn, 13))

	// https://en.wikipedia.org/wiki/Xorshift
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	return uint32(x)
}

// Uint32n returns a tread-safe, non-cryptographic pseudorandom uint32 in the range [0..maxN).
func Uint32n(maxN uint32) uint32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

func getRandomInt32() int32 {
	x := time.Now().UnixNano()
	return int32((x >> 32) ^ x)
}
