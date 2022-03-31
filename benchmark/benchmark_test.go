package benchmark

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
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

func BenchmarkGetWithIndex(b *testing.B) {
	b.StopTimer()

	ultraTable := perm()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.GetWithIdx("ID", i)
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

type T struct {
	Aa1 string `idx:"normal"`
	Ab1 string `idx:"normal"`
	Ac1 string `idx:"normal"`
	Ad1 string `idx:"normal"`
	Ae1 string `idx:"normal"`
	Af1 string `idx:"normal"`
	Ag1 string `idx:"normal"`
	Ah  string
	Ai  string
	Aj  string
	Ak  string
	Al  string
	Am  string
	An  string
	Ao  string
	Ap  string
	Aq  string
	Ar  string
	As  string
	At  string
	Au  string
	Av  string
	Aw  string
	Ax  string
	Ay  string
	Az  string
	Av1 string
	Aw1 string
	Ax1 string
	Ay1 string
	Az1 string
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

func BenchmarkCoverAdd(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.NewUltraTable()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.Add(T{
			Aa1: `hello world Aa1`,
			Ab1: `hello world Ab1`,
			Ac1: `hello world Ac1`,
			Ad1: `hello world Ad1`,
			Ae1: `hello world Ae1`,
			Af1: `hello world Af1`,
			Ag1: `hello world Ag1`,
			Ah:  `hello world Ah`,
			Ai:  `hello world Ai`,
			Aj:  `hello world Aj`,
			Ak:  `hello world Ak`,
			Al:  `hello world Al`,
			Am:  `hello world Am`,
			An:  `hello world An`,
			Ao:  `hello world Ao`,
			Ap:  `hello world Ap`,
			Aq:  `hello world Aq`,
			Ar:  `hello world Ar`,
			As:  `hello world As`,
			At:  `hello world At`,
			Au:  `hello world Au`,
			Av:  `hello world Av`,
			Aw:  `hello world Aw`,
			Ax:  `hello world Ax`,
			Ay:  `hello world Ay`,
			Az:  `hello world Az`,
			Av1: `hello world Av1`,
			Aw1: `hello world Aw1`,
			Ax1: `hello world Ax1`,
			Ay1: `hello world Ay1`,
			Az1: `hello world Az1`,
		})
	}
}

func BenchmarkCoverUpdate(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.NewUltraTable()
	for i := 1; i < 1000000; i++ {
		ultraTable.Add(Order{
			ID:        i,
			Account:   `a`,
			StockCode: `700`,
			Currency:  `HKD`,
			Amount:    100,
		})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ultraTable.UpdateWithIdx(`ID`, i, Order{
			ID:        i + 1000000,
			Account:   `a1`,
			StockCode: `800`,
			Currency:  `USD`,
			Amount:    100,
		})
	}
}

func BenchmarkCoverGet(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.NewUltraTable()
	for i := 1; i < 1000000; i++ {
		rand.Seed(time.Now().UnixNano())
		f := fmt.Sprintf(`%v`, rand.Intn(10000))
		ultraTable.Add(T{
			Aa1: `hello world ` + f,
			Ab1: `hello world ` + f,
			Ac1: `hello world ` + f,
			Ad1: `hello world ` + f,
			Ae1: `hello world ` + f,
			Af1: `hello world ` + f,
			Ag1: `hello world ` + f,
			Ah:  `hello world ` + f,
			Ai:  `hello world ` + f,
			Aj:  `hello world ` + f,
			Ak:  `hello world ` + f,
			Al:  `hello world ` + f,
			Am:  `hello world ` + f,
			An:  `hello world ` + f,
			Ao:  `hello world ` + f,
			Ap:  `hello world ` + f,
			Aq:  `hello world ` + f,
			Ar:  `hello world ` + f,
			As:  `hello world ` + f,
			At:  `hello world ` + f,
			Au:  `hello world ` + f,
			Av:  `hello world ` + f,
			Aw:  `hello world ` + f,
			Ax:  `hello world ` + f,
			Ay:  `hello world ` + f,
			Az:  `hello world ` + f,
			Av1: `hello world ` + f,
			Aw1: `hello world ` + f,
			Ax1: `hello world ` + f,
			Ay1: `hello world ` + f,
			Az1: `hello world ` + f,
		})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r, err := ultraTable.GetWithIdx(`Aa1`, `hello world `+`100`)
		if err != nil {
			b.Fatal(err)
		}
		if len(r) == 0 {
			b.Fail()
		}
	}
}

func BenchmarkConcurrent(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(30)

		ultraTable := ultra_table.NewUltraTable()
		for i := 0; i < 10; i++ {
			go func() {
				ultraTable.Add(Order{
					ID:        i,
					Account:   `a`,
					StockCode: `700`,
					Currency:  `HKD`,
					Amount:    100,
				})
				waitGroup.Done()
			}()
		}
		for i := 0; i < 10; i++ {
			go func() {
				ultraTable.UpdateWithIdx(`ID`, i, Order{
					ID:        i,
					Account:   `a1`,
					StockCode: `800`,
					Currency:  `USD`,
					Amount:    100,
				})
				waitGroup.Done()
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				ultraTable.GetWithIdx("ID", i)
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
		if ultraTable.Len() != 10 {
			b.Fail()
		}
	}
}

type IdempotentResponseBody struct {
	IdempotentKey string `idx:"normal"`
	TranctionTime int64  `idx:"normal"`
	IsErr         bool
	Body          []byte
}

func BenchmarkIdempotentCacheSize1(b *testing.B) {
	ultraTable := ultra_table.NewUltraTable()
	for i := 0; i < 200000; i++ {
		rand.Seed(time.Now().UnixNano())
		ultraTable.Add(IdempotentResponseBody{
			IdempotentKey: uuid.New().String(),
			TranctionTime: time.Now().Truncate(time.Hour * 24).Unix(),
			IsErr:         false,
			Body:          []byte("hello"),
		})
	}
	expired_time := time.Now()
	ultraTable.Remove(func(i interface{}) bool {
		tranctionTime := time.Unix(i.(IdempotentResponseBody).TranctionTime, 0)
		return expired_time.After(tranctionTime) || expired_time.Equal(tranctionTime)
	})
	if ultraTable.Len() != 0 {
		b.Fail()
	}
}

func BenchmarkIdempotentCacheSize2(b *testing.B) {
	ultraTable := ultra_table.NewUltraTable()
	for i := 0; i < 200000; i++ {
		rand.Seed(time.Now().UnixNano())
		ultraTable.Add(IdempotentResponseBody{
			IdempotentKey: uuid.New().String(),
			TranctionTime: time.Now().UnixMicro(),
			IsErr:         false,
			Body:          []byte("hello"),
		})
	}
	if ultraTable.Len() != 200000 {
		b.Fail()
	}
	expired_time := time.Now()
	ultraTable.Remove(func(i interface{}) bool {
		tranctionTime := time.UnixMicro(i.(IdempotentResponseBody).TranctionTime)
		return expired_time.After(tranctionTime) || expired_time.Equal(tranctionTime)
	})
	if ultraTable.Len() != 0 {
		b.Fail()
	}
	for i := 0; i < 200000; i++ {
		rand.Seed(time.Now().UnixNano())
		ultraTable.Add(IdempotentResponseBody{
			IdempotentKey: uuid.New().String(),
			TranctionTime: time.Now().UnixMicro(),
			IsErr:         false,
			Body:          []byte("hello"),
		})
	}
	if ultraTable.Len() != 200000 {
		b.Fail()
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkAddHasIndex-12                  1000000              1496 ns/op             477 B/op         11 allocs/op
// BenchmarkAdd-12                          3590232               368 ns/op             198 B/op          6 allocs/op
// BenchmarkGetWithIndex-12                19127990                63.6 ns/op             0 B/op          0 allocs/op
// BenchmarkRemoveWithIndex-12             16600911                67.3 ns/op             2 B/op          0 allocs/op
// BenchmarkAddRemove-12                     635726              2448 ns/op             750 B/op         26 allocs/op
// BenchmarkUpdateWithIdx-12                2637488               461 ns/op             138 B/op          3 allocs/op
// BenchmarkGetWithIdxIntersection-12       3900884               297 ns/op              17 B/op          1 allocs/op
// BenchmarkCoverAdd-12                      847106              1318 ns/op             711 B/op          8 allocs/op
// BenchmarkCoverUpdate-12                   594104              2488 ns/op             650 B/op         25 allocs/op
// BenchmarkCoverGet-12                      381310              2820 ns/op            1968 B/op         18 allocs/op
