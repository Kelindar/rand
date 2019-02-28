[![GoDoc](https://godoc.org/github.com/kelindar/rand?status.svg)](http://godoc.org/github.com/kelindar/rand)
[![Go Report](https://goreportcard.com/badge/github.com/kelindar/rand)](https://goreportcard.com/report/github.com/kelindar/rand)


# rand

An improved fast pseudorandom number generator based on https://github.com/valyala/fastrand.

# Features

- Optimized for speed.
- Performance scales on multiple CPUs.


# Benchmark results


```
$ GOMAXPROCS=1 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
pkg: github.com/kelindar/rand
BenchmarkUint32n                200000000               8.03 ns/op
BenchmarkMathRandInt31n         50000000                25.1 ns/op
```

```
$ GOMAXPROCS=2 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
pkg: github.com/kelindar/rand
BenchmarkUint32n-2              100000000               13.4 ns/op
BenchmarkMathRandInt31n-2       50000000                29.7 ns/op
```

```
$ GOMAXPROCS=4 go test -bench=. github.com/kelindar/rand
goos: windows
goarch: amd64
BenchmarkUint32n-4              100000000               16.0 ns/op
BenchmarkMathRandInt31n-4       30000000                46.1 ns/op
```

As you can see, [fastrand.Uint32n](https://godoc.org/github.com/kelindar/rand#Uint32n)
scales on multiple CPUs, while [rand.Int31n](https://golang.org/pkg/math/rand/#Int31n)
doesn't scale. Their performance is comparable on `GOMAXPROCS=1`,
but `rand.Uint32n` runs 2-3x faster than `rand.Int31n` on a multi-core system.
