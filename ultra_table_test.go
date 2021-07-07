package ultra_table

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUslice(t *testing.T) {
	Convey("Uslice", t, func() {
		Convey("Not Have Index", func() {
			type Order struct {
				ID        string
				Account   string
				StockCode string
				Currency  string
				Amount    float64
			}
			order := Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    55000,
			}
			Uslice := NewUltraTable()
			Uslice.Add(order)

			results := Uslice.Get(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(len(results), ShouldEqual, 1)
			_, err := Uslice.GetWithIdx("ID", "order_1")
			So(err, ShouldEqual, RecordNotFound)

			i := Uslice.Remove(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(i, ShouldEqual, 1)
		})
		Convey("Have Index", func() {
			type Order struct {
				ID        string `idx:"normal"`
				Account   string `idx:"normal"`
				StockCode string `idx:"normal"`
				Currency  string
				Amount    float64
			}
			order := Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    55000,
			}
			Uslice := NewUltraTable()
			Uslice.Add(order)

			results := Uslice.Get(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(len(results), ShouldEqual, 1)
			results, err := Uslice.GetWithIdx("ID", "order_1")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).Account == "1001"
			})
			So(len(results), ShouldEqual, 1)
			results, err = Uslice.GetWithIdx("Account", "1001")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).StockCode == "700"
			})
			So(len(results), ShouldEqual, 1)
			results, err = Uslice.GetWithIdx("StockCode", "700")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			i := Uslice.Remove(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(i, ShouldEqual, 1)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(len(results), ShouldEqual, 0)
			results, err = Uslice.GetWithIdx("ID", "order_1")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).Account == "1001"
			})
			So(len(results), ShouldEqual, 0)
			results, err = Uslice.GetWithIdx("Account", "1001")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).StockCode == "700"
			})
			So(len(results), ShouldEqual, 0)
			results, err = Uslice.GetWithIdx("StockCode", "700")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)
		})
		Convey("Have Index GetWithIdxIntersection", func() {
			type Order struct {
				ID        string `idx:"normal"`
				Account   string `idx:"normal"`
				StockCode string `idx:"normal"`
				Currency  string `idx:"normal"`
				Amount    float64
			}

			Uslice := NewUltraTable()
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1002",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1002",
				StockCode: "800",
				Currency:  "USD",
				Amount:    55000,
			})

			list, err := Uslice.GetWithIdxIntersection(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "700",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			list, err = Uslice.GetWithIdxIntersection(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 2)

			list, err = Uslice.GetWithIdxIntersection(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
				"Currency":  "SGD",
			})
			So(err, ShouldNotBeNil)
			So(len(list), ShouldEqual, 0)

			list, err = Uslice.GetWithIdxIntersection(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
				"Currency":  "USD",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 0)

			list, err = Uslice.GetWithIdxIntersection(map[string]interface{}{
				"Account":   "1002",
				"StockCode": "800",
				"Currency":  "USD",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)
		})
		Convey("Have Index GetWithIdxAggregate", func() {
			type Order struct {
				ID        string `idx:"normal"`
				Account   string `idx:"normal"`
				StockCode string `idx:"normal"`
				Currency  string `idx:"normal"`
				Amount    float64
			}

			Uslice := NewUltraTable()
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1001",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1002",
				StockCode: "800",
				Currency:  "HKD",
				Amount:    55000,
			})
			Uslice.Add(Order{
				ID:        "order_1",
				Account:   "1002",
				StockCode: "800",
				Currency:  "USD",
				Amount:    55000,
			})

			list, err := Uslice.GetWithIdxAggregate(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "700",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 3)

			list, err = Uslice.GetWithIdxAggregate(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 5)

			list, err = Uslice.GetWithIdxAggregate(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
				"Currency":  "SGD",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 5)

			list, err = Uslice.GetWithIdxAggregate(map[string]interface{}{
				"Account":   "1001",
				"StockCode": "800",
				"Currency":  "USD",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 5)

			list, err = Uslice.GetWithIdxAggregate(map[string]interface{}{
				"Account":   "1002",
				"StockCode": "800",
				"Currency":  "USD",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 4)

			type student struct {
				Name  string `idx:"normal"`
				Age   int    `idx:"normal"`
				Class string `idx:"normal"`
			}

			t := NewUltraTable()
			t.Add(student{
				Name:  `a`,
				Age:   18,
				Class: `1`,
			})
			t.Add(student{
				Name:  `b`,
				Age:   18,
				Class: `1`,
			})
			t.Add(student{
				Name:  `c`,
				Age:   20,
				Class: `2`,
			})
			t.Add(student{
				Name:  `d`,
				Age:   17,
				Class: `3`,
			})
			list, err = t.GetWithIdxAggregate(map[string]interface{}{
				"Name": "a",
				"Age":  18,
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 2)

			list, err = t.GetWithIdxAggregate(map[string]interface{}{
				"Class": "1",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 2)

			list, err = t.GetWithIdxAggregate(map[string]interface{}{
				"Name": "d",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			list, err = t.GetWithIdxAggregate(map[string]interface{}{
				"Age":   17,
				"Class": "0",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			list, err = t.GetWithIdxAggregate(map[string]interface{}{
				"Age":   17,
				"Class": "1",
			})
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 3)
		})
	})

}

func Test_Clear(t *testing.T) {
	Convey("Clear", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()
		for i := 0; i < 10000; i++ {
			ultraTable.Add(Order{
				ID:        fmt.Sprint(i),
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    500.1,
			})
		}
		So(ultraTable.Len(), ShouldEqual, 10000)
		So(ultraTable.Cap(), ShouldEqual, 10000)
		ultraTable.Clear()
		So(ultraTable.Len(), ShouldEqual, 0)
		So(ultraTable.Cap(), ShouldEqual, 0)
	})
}

func Test_Remove(t *testing.T) {
	Convey("Remove", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
			At        time.Time
		}
		Convey("Remove-1", func() {
			ultraTable := NewUltraTable()
			ultraTable.Add(Order{
				ID:        `1`,
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    100,
			})
			ultraTable.Add(Order{
				ID:        `1`,
				Account:   "1001",
				StockCode: "9988",
				Currency:  "HKD",
				Amount:    100,
			})
			So(ultraTable.RemoveWithIdx(`ID`, `1`), ShouldEqual, 2)
			So(ultraTable.Len(), ShouldEqual, 0)
			So(ultraTable.RemoveWithIdx(`ID`, `1`), ShouldEqual, 0)
			So(ultraTable.RemoveWithIdx(`Account`, `1001`), ShouldEqual, 0)
			So(ultraTable.RemoveWithIdx(`StockCode`, `700`), ShouldEqual, 0)
			So(ultraTable.Len(), ShouldEqual, 0)

			ultraTable.Add(Order{
				ID:        `2`,
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    100,
			})
			So(ultraTable.Len(), ShouldEqual, 1)
			ultraTable.Add(Order{
				ID:        `3`,
				Account:   "1001",
				StockCode: "700",
				Currency:  "HKD",
				Amount:    100,
			})
			So(ultraTable.Len(), ShouldEqual, 2)
			So(ultraTable.RemoveWithIdx(`Account`, `1001`), ShouldEqual, 2)
			So(ultraTable.RemoveWithIdx(`StockCode`, `700`), ShouldEqual, 0)
		})

		Convey("Remove-2", func() {
			ultraTable := NewUltraTable()
			for i := 0; i < 1000; i++ {
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: "700",
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 1000)
			for i := 0; i < 1000; i++ {
				items := ultraTable.Get(func(item interface{}) bool {
					return item.(Order).ID == fmt.Sprint(i)
				})
				So(items[0].(Order).Amount, ShouldEqual, i)

				isFound := ultraTable.Has(func(item interface{}) bool {
					return item.(Order).ID == fmt.Sprint(i)
				})
				So(isFound, ShouldBeTrue)
			}
			for i := 0; i < 1000; i++ {
				items, err := ultraTable.GetWithIdx("ID", fmt.Sprint(i))
				So(err, ShouldBeNil)
				So(items[0].(Order).Amount, ShouldEqual, i)

				isFound := ultraTable.HasWithIdx("ID", fmt.Sprint(i))
				So(isFound, ShouldBeTrue)
			}

			items := ultraTable.GetAll()
			So(len(items), ShouldEqual, 1000)

			for i := 0; i < 500; i++ {
				count := ultraTable.Remove(func(item interface{}) bool {
					return item.(Order).ID == fmt.Sprint(i)
				})
				So(count, ShouldEqual, 1)
			}
			So(ultraTable.Len(), ShouldEqual, 500)
			So(ultraTable.Cap(), ShouldEqual, 1000)
			for i := 500; i < 1000; i++ {
				count := ultraTable.RemoveWithIdx("ID", fmt.Sprint(i))
				So(count, ShouldEqual, 1)
			}
			So(ultraTable.Len(), ShouldEqual, 0)
			So(ultraTable.Cap(), ShouldEqual, 1000)

			items = ultraTable.GetAll()
			So(len(items), ShouldEqual, 0)
		})
		Convey("Remove-3", func() {
			ultraTable := NewUltraTable()

			for i := 0; i < 500; i++ {
				rand.Seed(time.Now().UnixNano())
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: fmt.Sprint(rand.Intn(1000)),
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}

			for i := 0; i < 10; i++ {
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: "00001",
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}

			for i := 0; i < 500; i++ {
				rand.Seed(time.Now().UnixNano())
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: fmt.Sprint(rand.Intn(1000)),
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 1010)
			So(ultraTable.Cap(), ShouldEqual, 1010)

			list, err := ultraTable.GetWithIdx("StockCode", "00001")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 10)

			count := ultraTable.RemoveWithIdx("StockCode", "00001")
			So(count, ShouldEqual, 10)
			So(ultraTable.Len(), ShouldEqual, 1000)
			So(ultraTable.Cap(), ShouldEqual, 1010)

			for i := 0; i < 10; i++ {
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: "00001",
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 1010)
			So(ultraTable.Cap(), ShouldEqual, 1010)

			for i := 0; i < 10; i++ {
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: "00001",
					Currency:  "HKD",
					Amount:    float64(i),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 1020)
			So(ultraTable.Cap(), ShouldEqual, 1020)
		})
		Convey("Remove-4", func() {
			ultraTable := NewUltraTable()

			for i := 0; i < 500; i++ {
				rand.Seed(time.Now().UnixNano())
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: fmt.Sprint(rand.Intn(1000)),
					Currency:  "HKD",
					Amount:    float64(i),
					At:        time.Now().Add(-time.Hour),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 500)
			So(ultraTable.Cap(), ShouldEqual, 500)
			ultraTable.Remove(func(i interface{}) bool {
				return i.(Order).At.Before(time.Now())
			})
			So(ultraTable.Len(), ShouldEqual, 0)
			So(ultraTable.Cap(), ShouldEqual, 500)

			for i := 0; i < 500; i++ {
				rand.Seed(time.Now().UnixNano())
				ultraTable.Add(Order{
					ID:        fmt.Sprint(i),
					Account:   "1001",
					StockCode: fmt.Sprint(rand.Intn(1000)),
					Currency:  "HKD",
					Amount:    float64(i),
					At:        time.Now().Add(-time.Hour),
				})
			}
			So(ultraTable.Len(), ShouldEqual, 500)
			So(ultraTable.Cap(), ShouldEqual, 500)
		})
	})
}

func Test_Update(t *testing.T) {
	Convey("Update", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()
		ultraTable.Add(Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})

		orders, err := ultraTable.GetWithIdx("ID", "order_1")
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.1)

		count := ultraTable.UpdateWithIdx("ID", "order_1", Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.2,
		})
		So(count, ShouldEqual, 1)

		orders, err = ultraTable.GetWithIdx("ID", "order_1")
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.2)

		ultraTable.Add(Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})
		ultraTable.Add(Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})
		ultraTable.Add(Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})

		count = ultraTable.UpdateWithIdx("ID", "order_2", Order{
			ID:        `order_3`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.2,
		})
		So(count, ShouldEqual, 3)

		orders, err = ultraTable.GetWithIdx("ID", "order_3")
		So(err, ShouldBeNil)
		So(len(orders), ShouldEqual, 3)
		So(orders[0].(Order).Amount, ShouldEqual, 500.2)
		So(orders[1].(Order).Amount, ShouldEqual, 500.2)
		So(orders[2].(Order).Amount, ShouldEqual, 500.2)

		count = ultraTable.UpdateWithIdx("StockCode", "700", Order{
			ID:        `order_3`,
			Account:   "1001",
			StockCode: "800",
			Currency:  "HKD",
			Amount:    500.2,
		})
		So(count, ShouldEqual, 4)

	})
}

func Test_Type(t *testing.T) {
	Convey("type", t, func() {

		Convey(`case-1`, func() {
			type TypeStruct struct {
				A string     `idx:"normal"`
				B int        `idx:"normal"`
				C int8       `idx:"normal"`
				D int16      `idx:"normal"`
				E int32      `idx:"normal"`
				F int64      `idx:"normal"`
				G uint       `idx:"normal"`
				H uint8      `idx:"normal"`
				I uint16     `idx:"normal"`
				J uint32     `idx:"normal"`
				K uint64     `idx:"normal"`
				L float32    `idx:"normal"`
				M float64    `idx:"normal"`
				N complex64  `idx:"normal"`
				O complex128 `idx:"normal"`
				P byte       `idx:"normal"`
				Q rune       `idx:"normal"`
			}
			ultraTable := NewUltraTable()
			for i := 0; i < 10; i++ {
				err := ultraTable.Add(TypeStruct{
					A: `test`,
					B: int(i),
					C: int8(i),
					D: int16(i),
					E: int32(i),
					F: int64(i),
					G: uint(i),
					H: uint8(i),
					I: uint16(i),
					J: uint32(i),
					K: uint64(i),
					L: float32(i),
					M: float64(i),
					N: 0,
					O: 0,
					P: byte(i),
					Q: rune(i),
				})
				So(err, ShouldBeNil)
			}
		})

		Convey(`case-2`, func() {
			type TypeStruct struct {
				A string      `idx:"normal"`
				B int         `idx:"normal"`
				C int8        `idx:"normal"`
				D int16       `idx:"normal"`
				E int32       `idx:"normal"`
				F int64       `idx:"normal"`
				G uint        `idx:"normal"`
				H uint8       `idx:"normal"`
				I uint16      `idx:"normal"`
				J uint32      `idx:"normal"`
				K uint64      `idx:"normal"`
				L float32     `idx:"normal"`
				M float64     `idx:"normal"`
				N complex64   `idx:"normal"`
				O complex128  `idx:"normal"`
				P byte        `idx:"normal"`
				Q rune        `idx:"normal"`
				X interface{} `idx:"normal"`
			}
			ultraTable := NewUltraTable()
			for i := 0; i < 10; i++ {
				err := ultraTable.Add(TypeStruct{
					A: `test`,
					B: int(i),
					C: int8(i),
					D: int16(i),
					E: int32(i),
					F: int64(i),
					G: uint(i),
					H: uint8(i),
					I: uint16(i),
					J: uint32(i),
					K: uint64(i),
					L: float32(i),
					M: float64(i),
					N: 0,
					O: 0,
					P: byte(i),
					Q: rune(i),
				})
				So(err, ShouldNotBeNil)
			}
		})

		Convey(`case-3`, func() {
			type TypeStruct struct {
				A string     `idx:"normal"`
				B int        `idx:"normal"`
				C int8       `idx:"normal"`
				D int16      `idx:"normal"`
				E int32      `idx:"normal"`
				F int64      `idx:"normal"`
				G uint       `idx:"normal"`
				H uint8      `idx:"normal"`
				I uint16     `idx:"normal"`
				J uint32     `idx:"normal"`
				K uint64     `idx:"normal"`
				L float32    `idx:"normal"`
				M float64    `idx:"normal"`
				N complex64  `idx:"normal"`
				O complex128 `idx:"normal"`
				P byte       `idx:"normal"`
				Q rune       `idx:"normal"`
				X []string   `idx:"normal"`
			}
			ultraTable := NewUltraTable()
			for i := 0; i < 10; i++ {
				err := ultraTable.Add(TypeStruct{
					A: `test`,
					B: int(i),
					C: int8(i),
					D: int16(i),
					E: int32(i),
					F: int64(i),
					G: uint(i),
					H: uint8(i),
					I: uint16(i),
					J: uint32(i),
					K: uint64(i),
					L: float32(i),
					M: float64(i),
					N: 0,
					O: 0,
					P: byte(i),
					Q: rune(i),
				})
				So(err, ShouldNotBeNil)
			}
		})

	})
}

func Test_SaveWithIdx(t *testing.T) {
	Convey("SaveWithIdx", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()

		count := ultraTable.SaveWithIdx(`ID`, `order_1`, Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})

		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 1)

		orders, err := ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.1)

		count = ultraTable.SaveWithIdx(`ID`, `order_1`, Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.2,
		})

		So(count, ShouldEqual, 1)

		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.2)
		So(ultraTable.Len(), ShouldEqual, 1)

		count = ultraTable.SaveWithIdx(`Account`, `1001`, Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.3,
		})

		So(count, ShouldEqual, 1)

		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.3)
		So(ultraTable.Len(), ShouldEqual, 1)

		count = ultraTable.SaveWithIdx(`Account`, `1002`, Order{
			ID:        `order_1`,
			Account:   "1002",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.5,
		})

		So(count, ShouldEqual, 1)

		orders, err = ultraTable.GetWithIdx(`Account`, `1002`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.5)
		So(ultraTable.Len(), ShouldEqual, 2)
	})
}

func Test_SaveWithIdxAggregate(t *testing.T) {
	Convey("SaveWithIdx", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()

		count := ultraTable.SaveWithIdxAggregate(map[string]interface{}{`ID`: `order_1`}, Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})

		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 1)

		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`ID`: `order_2`}, Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "800",
			Currency:  "HKD",
			Amount:    500.1,
		})

		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 2)

		list, err := ultraTable.GetWithIdxIntersection(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`})
		So(err, ShouldBeNil)
		So(len(list), ShouldEqual, 1)

		list, err = ultraTable.GetWithIdxAggregate(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`})
		So(err, ShouldBeNil)
		So(len(list), ShouldEqual, 2)

		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`}, Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.2,
		})
		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 2)

		for idx, v := range ultraTable.GetAll() {
			if idx == 0 {
				So(v.(Order).Amount, ShouldEqual, 500.2)
			}
			if idx == 1 {
				So(v.(Order).Amount, ShouldEqual, 500.1)
			}
		}
		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`}, Order{
			ID:        `order_2`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.3,
		})
		So(count, ShouldEqual, 2)
		So(ultraTable.Len(), ShouldEqual, 2)

		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`Account`: `1003`, `StockCode`: `900`}, Order{
			ID:        `order_3`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.4,
		})

		for idx, v := range ultraTable.GetAll() {
			if idx == 0 {
				So(v.(Order).ID, ShouldEqual, `order_2`)
				So(v.(Order).Amount, ShouldEqual, 500.3)
			}
			if idx == 1 {
				So(v.(Order).ID, ShouldEqual, `order_2`)
				So(v.(Order).Amount, ShouldEqual, 500.3)
			}
			if idx == 2 {
				So(v.(Order).ID, ShouldEqual, `order_3`)
				So(v.(Order).Amount, ShouldEqual, 500.4)
			}
		}

		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1005`, `StockCode`: `900`}, Order{
			ID:        `order_4`,
			Account:   "1005",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.5,
		})
		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 4)

		for idx, v := range ultraTable.GetAll() {
			if idx == 0 {
				So(v.(Order).ID, ShouldEqual, `order_2`)
				So(v.(Order).Amount, ShouldEqual, 500.3)
			}
			if idx == 1 {
				So(v.(Order).ID, ShouldEqual, `order_2`)
				So(v.(Order).Amount, ShouldEqual, 500.3)
			}
			if idx == 2 {
				So(v.(Order).ID, ShouldEqual, `order_3`)
				So(v.(Order).Amount, ShouldEqual, 500.4)
			}
			if idx == 3 {
				So(v.(Order).ID, ShouldEqual, `order_4`)
				So(v.(Order).Amount, ShouldEqual, 500.5)
			}
		}

		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1005`, `StockCode`: `1000`}, Order{
			ID:        `order_5`,
			Account:   "1006",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.5,
		})
		So(count, ShouldEqual, 1)
		So(ultraTable.Len(), ShouldEqual, 5)

	})
}

func Test_Kind(t *testing.T) {
	Convey("struct", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()
		ultraTable.Add(Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})
		orders, err := ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.1)

		order := orders[0].(Order)
		order.Amount = 500.2

		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.1)

		count := ultraTable.UpdateWithIdx("Account", "1001", Order{
			ID:        `order_3`,
			Account:   "1001",
			StockCode: "800",
			Currency:  "HKD",
			Amount:    500.2,
		})
		So(count, ShouldEqual, 1)

		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
		So(err, ShouldBeNil)
		So(orders[0].(Order).Amount, ShouldEqual, 500.2)
	})

	Convey("ptr", t, func() {
		type Order struct {
			ID        string `idx:"normal"`
			Account   string `idx:"normal"`
			StockCode string `idx:"normal"`
			Currency  string
			Amount    float64
		}
		ultraTable := NewUltraTable()
		err := ultraTable.Add(&Order{
			ID:        `order_1`,
			Account:   "1001",
			StockCode: "700",
			Currency:  "HKD",
			Amount:    500.1,
		})
		So(err, ShouldEqual, OnlySupportStruct)

	})
}
