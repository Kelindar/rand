package rand

import (
	"math/rand"
	"sync/atomic"
	"testing"
)

// BenchSink prevents the compiler from optimizing away benchmark loops.
var BenchSink uint64

func BenchmarkRand(b *testing.B) {
	b.Run("Uint32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint32(0)
			for pb.Next() {
				s += Uint32()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint32n", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint32(0)
			for pb.Next() {
				s += Uint32n(1e6)
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Float32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := float32(0)
			for pb.Next() {
				s += Float32()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint64", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint64(0)
			for pb.Next() {
				s += Uint64()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint64n", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint64(0)
			for pb.Next() {
				s += Uint64n(1e6)
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Float64", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := float64(0)
			for pb.Next() {
				s += Float64()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})
}

func BenchmarkMathRandInt31n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += uint32(rand.Int31n(1e6))
		}
		atomic.AddUint64(&BenchSink, uint64(s))
	})
}
