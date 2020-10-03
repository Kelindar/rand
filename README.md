[![GoDoc](https://godoc.org/github.com/kelindar/rand?status.svg)](http://godoc.org/github.com/kelindar/rand)
[![Go Report](https://goreportcard.com/badge/github.com/kelindar/rand)](https://goreportcard.com/report/github.com/kelindar/rand)


# rand

A simple thread-safe API for a Go implementation of Melissa O'Neill's excellent PCG pseudorandom number generator, which is well-studied, excellent, and fast both to create and in execution.

This repository contains a fork of [MichaelTJones/pcg](https://github.com/MichaelTJones/pcg) code and adds a pooling of random number generators.

# Features

- Optimized for speed.
- Performance scales on multiple CPUs.


# Benchmark results


```
BenchmarkRand/Uint32-8         	566756457	         2.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/Uint32n-8        	478460296	         2.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/Float32-8        	406576479	         2.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/Uint64-8         	469784265	         2.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/Uint64n-8        	273879046	         4.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/Float64-8        	396292614	         3.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkMathRandInt31n-8   	29865824	        40.2 ns/op	       0 B/op	       0 allocs/op
```
