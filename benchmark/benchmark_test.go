package benchmark

import (
	"testing"
	"ultra_table"
)

type Order struct {
	ID        int    `index:"id"`
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
			ID:        i,
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
		ID        int
		Account   string
		StockCode string
		Currency  string
		Amount    float64
	}
	ultraTable := ultra_table.NewUltraTable()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ultraTable.Add(Order{
			ID:        i,
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
			ID:        i,
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
			return i.(Order).ID == i
		})
	}
}

func BenchmarkGetWithIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.GetWithIdx("id", i)
	}
}

func BenchmarkRemove(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.Remove(func(i interface{}) bool {
			return i.(Order).ID == i
		})
	}

}

func BenchmarkRemoveWithIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.RemoveWithIdx("id", i)
	}
}

func BenchmarkAddRemove(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.Add(Order{
			ID:        i,
			Account:   "1001",
			StockCode: "00001",
			Currency:  "HKD",
			Amount:    float64(i),
		})
		ultraTable.RemoveWithIdx("id", i)
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkAddHasIndex
// BenchmarkAddHasIndex-12           905406              1957 ns/op             603 B/op          8 allocs/op
// BenchmarkAdd
// BenchmarkAdd-12                  3691093               323 ns/op             195 B/op          6 allocs/op
// BenchmarkGet
// BenchmarkGet-12                     1904            630317 ns/op               0 B/op          0 allocs/op
// BenchmarkGetWithIndex
// BenchmarkGetWithIndex-12        13605790                92.0 ns/op             0 B/op          0 allocs/op
// BenchmarkRemove
// BenchmarkRemove-12                  1972            666773 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithIndex
// BenchmarkRemoveWithIndex-12     17592115                62.7 ns/op             0 B/op          0 allocs/op
// BenchmarkAddRemove
// BenchmarkAddRemove-12             460332              2305 ns/op             454 B/op         15 allocs/op
// PASS
