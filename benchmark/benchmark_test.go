package benchmark

import (
	"testing"

	ultra_table "github.com/longbridgeapp/ultra-table"
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

func BenchmarkGetWithIdxIntersection(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.GetWithIdxIntersection(map[string]interface{}{
			"id":         i,
			"account":    "1001",
			"stock_code": "00001",
		})
	}
}

// goarch: amd64
// BenchmarkAddHasIndex
// BenchmarkAddHasIndex-12                   865812              1770 ns/op             599 B/op          8 allocs/op
// BenchmarkAdd
// BenchmarkAdd-12                          3714140               321 ns/op             195 B/op          6 allocs/op
// BenchmarkGet
// BenchmarkGet-12                             1982            599333 ns/op               0 B/op          0 allocs/op
// BenchmarkGetWithIndex
// BenchmarkGetWithIndex-12                20277746                58.2 ns/op             0 B/op          0 allocs/op
// BenchmarkRemove
// BenchmarkRemove-12                          2148            546817 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithIndex
// BenchmarkRemoveWithIndex-12             16901769                64.6 ns/op             0 B/op          0 allocs/op
// BenchmarkAddRemove
// BenchmarkAddRemove-12                     436630              2326 ns/op             459 B/op         15 allocs/op
// BenchmarkUpdateWithIdx
// BenchmarkUpdateWithIdx-12                2973962               389 ns/op             103 B/op          2 allocs/op
// BenchmarkGetWithIdxIntersection
// BenchmarkGetWithIdxIntersection-12       5756240               200 ns/op              15 B/op          1 allocs/op
