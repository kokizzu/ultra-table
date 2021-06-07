package benchmark

import (
	"testing"
	ultra_table "ultra-table"
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

func BenchmarkUpdateWithIdx(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.UpdateWithIdx("id", i, Order{
			ID:        i + 1000000,
			Account:   "1002",
			StockCode: "00002",
			Currency:  "USD",
			Amount:    float64(i),
		})
	}
}

// BenchmarkAddHasIndex
// BenchmarkAddHasIndex-12           840850              1787 ns/op             481 B/op          8 allocs/op
// BenchmarkAdd
// BenchmarkAdd-12                  3619507               345 ns/op             197 B/op          6 allocs/op
// BenchmarkGet
// BenchmarkGet-12                     1663            765931 ns/op               0 B/op          0 allocs/op
// BenchmarkGetWithIndex
// BenchmarkGetWithIndex-12        18395773                80.6 ns/op             0 B/op          0 allocs/op
// BenchmarkRemove
// BenchmarkRemove-12                  1910            557853 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithIndex
// BenchmarkRemoveWithIndex-12     17457786                63.3 ns/op             0 B/op          0 allocs/op
// BenchmarkAddRemove
// BenchmarkAddRemove-12             464210              2489 ns/op             453 B/op         15 allocs/op
// BenchmarkUpdateWithIdx
// BenchmarkUpdateWithIdx-12        2600654               411 ns/op             109 B/op          2 allocs/op
// PASS
