package benchmark

import (
	"fmt"
	"testing"
	"ultra_table"
)

type Order struct {
	ID        string `index:"id"`
	Account   string `index:"account"`
	StockCode string `index:"stock_code"`
	Currency  string
	Amount    float64
}

func BenchmarkAddHasIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := ultra_table.NewUltraTable()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ultraTable.Add(Order{
			ID:        fmt.Sprint(i),
			Account:   "1001",
			StockCode: "00001",
			Currency:  "HKD",
			Amount:    float64(i),
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	type Order struct {
		ID        string
		Account   string
		StockCode string
		Currency  string
		Amount    float64
	}
	ultraTable := ultra_table.NewUltraTable()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ultraTable.Add(Order{
			ID:        fmt.Sprint(i),
			Account:   "1001",
			StockCode: "00001",
			Currency:  "HKD",
			Amount:    float64(i),
		})
	}
}

func perm() *ultra_table.UltraTable {
	ultraTable := ultra_table.NewUltraTable()

	for i := 0; i < 100000; i++ {
		ultraTable.Add(Order{
			ID:        fmt.Sprint(i),
			Account:   "1001",
			StockCode: "00001",
			Currency:  "HKD",
			Amount:    float64(i),
		})
	}
	return ultraTable
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.Get(func(i interface{}) bool {
			return i.(Order).ID == fmt.Sprint(i)
		})
	}
}

func BenchmarkGetWithIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.GetWithIdx("id", fmt.Sprint(i))
	}
}

func BenchmarkRemove(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.Remove(func(i interface{}) bool {
			return i.(Order).ID == fmt.Sprint(i)
		})
	}

}

func BenchmarkRemoveWithIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.RemoveWithIdx("id", fmt.Sprint(i))
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkAddHasIndex
// BenchmarkAddHasIndex-12           969079              1615 ns/op             535 B/op          9 allocs/op
// BenchmarkAdd
// BenchmarkAdd-12                  2730561               439 ns/op             235 B/op          7 allocs/op
// BenchmarkGet
// BenchmarkGet-12                       24          46857622 ns/op         3200236 B/op     100001 allocs/op
// BenchmarkGetWithIndex
// BenchmarkGetWithIndex-12         7500440               152 ns/op              16 B/op          2 allocs/op
// BenchmarkRemove
// BenchmarkRemove-12                    24          47166444 ns/op         3200223 B/op     100000 allocs/op
// BenchmarkRemoveWithIndex
// BenchmarkRemoveWithIndex-12      5563045               186 ns/op              32 B/op          3 allocs/op
