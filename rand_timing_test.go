package rand

import (
	"math/rand"
	"sync/atomic"
	"testing"
)

// BenchSink prevents the compiler from optimizing away benchmark loops.
var BenchSink uint32

func BenchmarkUint32n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += Uint32n(1e6)
		}
		atomic.AddUint32(&BenchSink, s)
	})
}

func BenchmarkMathRandInt31n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += uint32(rand.Int31n(1e6))
		}
		atomic.AddUint32(&BenchSink, s)
	})
}
