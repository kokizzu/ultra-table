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
BenchmarkAddHasIndex                     4793829              3639 ns/op             456 B/op         11 allocs/op
BenchmarkAdd                            25575007               457 ns/op             202 B/op          6 allocs/op
BenchmarkGetWithIndex                   206441649               58.4 ns/op             0 B/op          0 allocs/op
BenchmarkRemoveWithIndex                194092122               55.9 ns/op             0 B/op          0 allocs/op
BenchmarkAddRemove                       2764548              5311 ns/op             764 B/op         26 allocs/op
BenchmarkUpdateWithIdx                  13222945               830 ns/op             134 B/op          3 allocs/op
BenchmarkGetWithIdxIntersection         44541792               237 ns/op              12 B/op          1 allocs/op
BenchmarkCoverAdd                        7283416              1830 ns/op             718 B/op          8 allocs/op
BenchmarkCoverUpdate                     2718184              5297 ns/op             777 B/op         25 allocs/op
BenchmarkCoverGet                        2192914              5439 ns/op            2320 B/op         17 allocs/op
```