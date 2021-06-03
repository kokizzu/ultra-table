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
			_, err := Uslice.GetWithIdx("id", "order_1")
			So(err, ShouldEqual, RecordNotFound)

			i := Uslice.Remove(func(i interface{}) bool {
				return i.(Order).ID == "order_1"
			})
			So(i, ShouldEqual, 1)
		})
		Convey("Have Index", func() {
			type Order struct {
				ID        string `index:"id"`
				Account   string `index:"account"`
				StockCode string `index:"stock_code"`
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
			results, err := Uslice.GetWithIdx("id", "order_1")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).Account == "1001"
			})
			So(len(results), ShouldEqual, 1)
			results, err = Uslice.GetWithIdx("account", "1001")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).StockCode == "700"
			})
			So(len(results), ShouldEqual, 1)
			results, err = Uslice.GetWithIdx("stock_code", "700")
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
			results, err = Uslice.GetWithIdx("id", "order_1")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).Account == "1001"
			})
			So(len(results), ShouldEqual, 0)
			results, err = Uslice.GetWithIdx("account", "1001")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)

			results = Uslice.Get(func(i interface{}) bool {
				return i.(Order).StockCode == "700"
			})
			So(len(results), ShouldEqual, 0)
			results, err = Uslice.GetWithIdx("stock_code", "700")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)
		})
	})

}

func Test_Clear(t *testing.T) {
	Convey("Clear", t, func() {
		type Order struct {
			ID        string `index:"id"`
			Account   string `index:"account"`
			StockCode string `index:"stock_code"`
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
			ID        string `index:"id"`
			Account   string `index:"account"`
			StockCode string `index:"stock_code"`
			Currency  string
			Amount    float64
		}
		Convey("Remove-1", func() {
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
				items, err := ultraTable.GetWithIdx("id", fmt.Sprint(i))
				So(err, ShouldBeNil)
				So(items[0].(Order).Amount, ShouldEqual, i)

				isFound := ultraTable.HasWithIdx("id", fmt.Sprint(i))
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
				count := ultraTable.RemoveWithIdx("id", fmt.Sprint(i))
				So(count, ShouldEqual, 1)
			}
			So(ultraTable.Len(), ShouldEqual, 0)
			So(ultraTable.Cap(), ShouldEqual, 1000)

			items = ultraTable.GetAll()
			So(len(items), ShouldEqual, 0)
		})
		Convey("Remove-2", func() {
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

			list, err := ultraTable.GetWithIdx("stock_code", "00001")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 10)

			count := ultraTable.RemoveWithIdx("stock_code", "00001")
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
	})
}
