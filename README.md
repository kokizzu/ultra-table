## Overview
Ultra_table provides a possibility to quickly build an in-memory table, It has good performance and high scalability, The current version already supports serialization of gogoproto and easyjson, The most important thing is that he is based on the latest golang 1.18 implementation, which is the implementation of go generics.
We try to apply sdk to internal applications.
Requires Go 1.18 or newer.
### Getting started
To install Ultra_table, use `go get`:

```sh
$ go get github.com/longbridgeapp/ultra-table
```

### Example
```
package main

import (
	"log"

	ultra_table "github.com/longbridgeapp/ultra-table"
	"github.com/longbridgeapp/ultra-table/test_data/easyjson"
	"github.com/longbridgeapp/ultra-table/test_data/pb"
)

func main() {
	baseEasyjson() //serialization based easyjson
	basegogo()     //serialization based gogo protobuf
}

func baseEasyjson() {
	table := ultra_table.New[*easyjson.Person]()

	err := table.Add(&easyjson.Person{
		Name:     "jacky",
		Phone:    "+8613575468007",
		Age:      31,
		BirthDay: 19901111,
		Gender:   0,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = table.Add(&easyjson.Person{
		Name:     "rose",
		Phone:    "+8613575468008",
		Age:      31,
		BirthDay: 19901016,
		Gender:   1,
	})
	if err != nil {
		log.Fatalln("easyjson", err)
	}

	infos, err := table.GetWithIdx("Phone", "+8613575468007")
	if err != nil {
		log.Fatalln("easyjson", err)
	}
	for i := 0; i < len(infos); i++ {
		log.Printf("easyjson %+v \n", infos[i])
	}

	_, err = table.GetWithIdxIntersection(map[string]interface{}{
		"Age":  31,
		"Name": "rose",
	})
	log.Println("easyjson", err)
}

func basegogo() {
	table := ultra_table.New[*pb.Person]()

	err := table.Add(&pb.Person{
		Name:     "jacky",
		Phone:    "+8613575468007",
		Age:      31,
		BirthDay: 19901111,
		Gender:   pb.Gender_men,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = table.Add(&pb.Person{
		Name:     "rose",
		Phone:    "+8613575468008",
		Age:      31,
		BirthDay: 19901016,
		Gender:   pb.Gender_women,
	})
	if err != nil {
		log.Fatalln("gogo", err)
	}

	infos, err := table.GetWithIdx("Phone", "+8613575468007")
	if err != nil {
		log.Fatalln("gogo", err)
	}
	for i := 0; i < len(infos); i++ {
		log.Printf("gogo %+v \n", infos[i])
	}

	_, err = table.GetWithIdxIntersection(map[string]interface{}{
		"Age":  31,
		"Name": "rose",
	})
	log.Println("gogo", err)
}
```

### Benchmarks

```
GOMAXPROCS=1 go test -v -bench=. -benchmem benchmark_test.go -benchtime=10s
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAddWithGoGo                             3587053              4568 ns/op             713 B/op         16 allocs/op
BenchmarkAddWithEasyjson                         2884038              5259 ns/op             667 B/op         16 allocs/op
BenchmarkGetWithUniqueIndexWithGoGo             56969133               220.0 ns/op            29 B/op          3 allocs/op
BenchmarkGetWithUniqueIndexWithEasyjson         27537194               504.4 ns/op            29 B/op          3 allocs/op
BenchmarkGetWithNormalIndex                     25512718               416.2 ns/op           109 B/op          5 allocs/op
BenchmarkGetWithIdxIntersectionNotFound          9719488              1231 ns/op             304 B/op         13 allocs/op
BenchmarkGetWithIdxIntersection                  9484707              1474 ns/op             336 B/op         16 allocs/op
BenchmarkRemoveWithIndex                        60868005               191.7 ns/op            32 B/op          2 allocs/op
BenchmarkUpdateWithIndex                        16691726               710.2 ns/op           148 B/op          5 allocs/op
BenchmarkAddAndRemove                            4346674              2928 ns/op            1002 B/op         45 allocs/op
```