# ultra_table is an implementation of memory table for Go
=======

## Overview
Ultra_table is an zero copy of a memory table, which supports the automatic generation of multi-dimensional memory indexes based on the struct, so that index-based query, delete, and update achieve a relatively good performance.

We try to apply sdk to internal applications.
At the same time, I hope to receive more suggestions.

### Getting started
To install Ultra_table, use `go get`:

```sh
$ go get github.com/longbridgeapp/ultra-table
```

### Benchmark

```
$ GOMAXPROCS=1 go test -v -bench=. -benchmem benchmark_test.go -benchtime=10s
BenchmarkAddHasIndex                     5330475              2807 ns/op             524 B/op          6 allocs/op
BenchmarkAdd                            28591881               443 ns/op             192 B/op          6 allocs/op
BenchmarkGetWithIndex                   210208250               56.4 ns/op             0 B/op          0 allocs/op
BenchmarkRemoveWithIndex                200379544               62.2 ns/op             0 B/op          0 allocs/op
BenchmarkAddRemove                       3409944              3181 ns/op             492 B/op         11 allocs/op
BenchmarkUpdateWithIdx                  24096828               654 ns/op             104 B/op          2 allocs/op
BenchmarkGetWithIdxIntersection         49789596               244 ns/op              12 B/op          1 allocs/op
```