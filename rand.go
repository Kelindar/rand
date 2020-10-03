// Package rand implements fast pesudorandom number generator
// that should scale well on multi-CPU systems.
//
// Use crypto/rand instead of this package for generating
// cryptographically secure random numbers.
package rand

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"sync"
)

// Pool for the random number generators
var pool32 = sync.Pool{
	New: func() interface{} {
		return newRand32().Seed(seed(), seed())
	},
}

// Pool for the random number generators
var pool64 = sync.Pool{
	New: func() interface{} {
		return newRand64().Seed(seed(), seed(), seed(), seed())
	},
}

// seed generates a crypto-random seed. This is slow but is used to seed the random
// number generators themselves.
func seed() uint64 {
	buffer := make([]byte, 8)
	rand.Read(buffer)
	return binary.LittleEndian.Uint64(buffer[:8])
}

// ----------------------------------------------------------------------------------------------

// Bool returns, as a bool, a pseudo-random number
func Bool() bool {
	return Uint32n(2) == 0
}

// Uint32 returns a tread-safe, non-cryptographic pseudorandom uint32.
func Uint32() uint32 {
	rng := pool32.Get().(*rand32)
	defer pool32.Put(rng)
	return rng.Random()
}

// Uint32n returns a tread-safe, non-cryptographic pseudorandom uint32 in the range [0..maxN).
func Uint32n(maxN uint32) uint32 {
	rng := pool32.Get().(*rand32)
	defer pool32.Put(rng)
	return rng.Bounded(maxN)
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0)
func Float32() float32 {
	return math.Float32frombits(Uint32()) / (1 << 31)
}

// ----------------------------------------------------------------------------------------------

// Uint64 returns a tread-safe, non-cryptographic pseudorandom uint64.
func Uint64() uint64 {
	rng := pool64.Get().(*rand64)
	defer pool64.Put(rng)
	return rng.Random()
}

// Uint64n returns a tread-safe, non-cryptographic pseudorandom uint64 in the range [0..maxN).
func Uint64n(maxN uint64) uint64 {
	rng := pool64.Get().(*rand64)
	defer pool64.Put(rng)
	return rng.Bounded(maxN)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0)
func Float64() float64 {
	return math.Float64frombits(Uint64()) / (1 << 63)
}
