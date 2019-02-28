[![GoDoc](https://godoc.org/github.com/kelindar/rand?status.svg)](http://godoc.org/github.com/kelindar/rand)
[![Go Report](https://goreportcard.com/badge/github.com/kelindar/rand)](https://goreportcard.com/report/github.com/kelindar/rand)


# rand

An improved fast pseudorandom number generator based on https://github.com/valyala/fastrand.

# Features

- Optimized for speed.
- Performance scales on multiple CPUs.

# How does it work?

It abuses [sync.Pool](https://golang.org/pkg/sync/#Pool) for maintaining
"per-CPU" pseudorandom number generators.

TODO: firgure out how to use real per-CPU pseudorandom number generators.


# Benchmark results


```
$ GOMAXPROCS=1 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
pkg: github.com/kelindar/rand
BenchmarkUint32n                        50000000                24.3 ns/op
BenchmarkRNGUint32n                     500000000               3.66 ns/op
BenchmarkRNGUint32nWithLock             100000000               17.3 ns/op
BenchmarkRNGUint32nArray                100000000               25.0 ns/op
BenchmarkMathRandInt31n                 50000000                25.0 ns/op
BenchmarkMathRandRNGInt31n              200000000               10.0 ns/op
BenchmarkMathRandRNGInt31nWithLock      100000000               22.2 ns/op
BenchmarkMathRandRNGInt31nArray         30000000                45.6 ns/op
PASS
ok      github.com/kelindar/rand        17.219s
```

```
$ GOMAXPROCS=2 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
pkg: github.com/kelindar/rand
BenchmarkUint32n-2                      100000000               17.6 ns/op
BenchmarkRNGUint32n-2                   2000000000              1.85 ns/op
BenchmarkRNGUint32nWithLock-2           50000000                22.1 ns/op
BenchmarkRNGUint32nArray-2              30000000                38.6 ns/op
BenchmarkMathRandInt31n-2               50000000                29.8 ns/op
BenchmarkMathRandRNGInt31n-2            300000000               5.15 ns/op
BenchmarkMathRandRNGInt31nWithLock-2    50000000                30.0 ns/op
BenchmarkMathRandRNGInt31nArray-2       20000000                61.1 ns/op
PASS
ok      github.com/kelindar/rand        15.790s
```

```
$ GOMAXPROCS=4 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
pkg: github.com/kelindar/rand
BenchmarkUint32n-4                      200000000                7.58 ns/op
BenchmarkRNGUint32n-4                   2000000000               1.12 ns/op
BenchmarkRNGUint32nWithLock-4           50000000                36.0 ns/op
BenchmarkRNGUint32nArray-4              50000000                33.8 ns/op
BenchmarkMathRandInt31n-4               30000000                46.1 ns/op
BenchmarkMathRandRNGInt31n-4            1000000000               2.87 ns/op
BenchmarkMathRandRNGInt31nWithLock-4    30000000                43.4 ns/op
BenchmarkMathRandRNGInt31nArray-4       30000000                38.9 ns/op
PASS
ok      github.com/kelindar/rand        15.953s
```

As you can see, [fastrand.Uint32n](https://godoc.org/github.com/kelindar/rand#Uint32n)
scales on multiple CPUs, while [rand.Int31n](https://golang.org/pkg/math/rand/#Int31n)
doesn't scale. Their performance is comparable on `GOMAXPROCS=1`,
but `rand.Uint32n` runs 3x faster than `rand.Int31n` on `GOMAXPROCS=2`
and 10x faster than `rand.Int31n` on `GOMAXPROCS=4`.
