package benchmark

import (
	"testing"

	ultra_table "github.com/longbridgeapp/ultra-table"
)

type Order struct {
	ID        int    `idx:"normal"`
	Account   string `idx:"normal"`
	StockCode string `idx:"normal"`
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
		ultraTable.GetWithIdx("ID", i)
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
		ultraTable.RemoveWithIdx("ID", i)
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
		ultraTable.RemoveWithIdx("ID", i)
	}
}

func BenchmarkUpdateWithIdx(b *testing.B) {
	b.StopTimer()
	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.UpdateWithIdx("ID", i, Order{
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
			"ID":        i,
			"Account":   "1001",
			"StockCode": "00001",
		})
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkAddHasIndex
// BenchmarkAddHasIndex-12                  1000000              1721 ns/op             570 B/op          6 allocs/op
// BenchmarkAdd
// BenchmarkAdd-12                          3491673               337 ns/op             200 B/op          6 allocs/op
// BenchmarkGet
// BenchmarkGet-12                             1842            609910 ns/op               0 B/op          0 allocs/op
// BenchmarkGetWithIndex
// BenchmarkGetWithIndex-12                15409140                67.2 ns/op             0 B/op          0 allocs/op
// BenchmarkRemove
// BenchmarkRemove-12                          1665            755570 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithIndex
// BenchmarkRemoveWithIndex-12             17793783                61.1 ns/op             0 B/op          0 allocs/op
// BenchmarkAddRemove
// BenchmarkAddRemove-12                     669391              1659 ns/op             376 B/op          7 allocs/op
// BenchmarkUpdateWithIdx
// BenchmarkUpdateWithIdx-12                3056722               336 ns/op             112 B/op          1 allocs/op
// BenchmarkGetWithIdxIntersection
// BenchmarkGetWithIdxIntersection-12       5862406               208 ns/op              15 B/op          1 allocs/op
