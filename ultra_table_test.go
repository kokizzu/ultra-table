package ultra_table

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Add", t, func() {
		Convey("have not index", func() {

		})
		Convey("have not index", func() {

		})
	})
}

// func TestTable(t *testing.T) {
// 	Convey("Table", t, func() {
// 		Convey("Not Have Index", func() {
// 			type Order struct {
// 				ID        string
// 				Account   string
// 				StockCode string
// 				Currency  string
// 				Amount    float64
// 			}
// 			order := Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			}
// 			table := New[Order]()
// 			table.Add(order)

// 			results := table.Get(func(order Order) bool {
// 				return order.ID == "order_1"
// 			})
// 			So(len(results), ShouldEqual, 1)
// 			_, err := table.GetWithIdx("ID", "order_1")
// 			So(err, ShouldEqual, RecordNotFound)

// 			i := table.Remove(func(order Order) bool {
// 				return order.ID == "order_1"
// 			})
// 			So(i, ShouldEqual, 1)
// 		})
// 		Convey("Have Index", func() {
// 			type Order struct {
// 				ID        string `idx:"normal"`
// 				Account   string `idx:"normal"`
// 				StockCode string `idx:"normal"`
// 				Currency  string
// 				Amount    float64
// 			}
// 			order := Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			}
// 			table := New[Order]()
// 			table.Add(order)

// 			results := table.Get(func(order Order) bool {
// 				return order.ID == "order_1"
// 			})
// 			So(len(results), ShouldEqual, 1)
// 			results, err := table.GetWithIdx("ID", "order_1")
// 			So(err, ShouldBeNil)
// 			So(len(results), ShouldEqual, 1)

// 			results = table.Get(func(order Order) bool {
// 				return order.Account == "1001"
// 			})
// 			So(len(results), ShouldEqual, 1)
// 			results, err = table.GetWithIdx("Account", "1001")
// 			So(err, ShouldBeNil)
// 			So(len(results), ShouldEqual, 1)

// 			results = table.Get(func(order Order) bool {
// 				return order.StockCode == "700"
// 			})
// 			So(len(results), ShouldEqual, 1)
// 			results, err = table.GetWithIdx("StockCode", "700")
// 			So(err, ShouldBeNil)
// 			So(len(results), ShouldEqual, 1)

// 			i := table.Remove(func(order Order) bool {
// 				return order.ID == "order_1"
// 			})
// 			So(i, ShouldEqual, 1)

// 			results = table.Get(func(order Order) bool {
// 				return order.ID == "order_1"
// 			})
// 			So(len(results), ShouldEqual, 0)
// 			results, err = table.GetWithIdx("ID", "order_1")
// 			So(err, ShouldNotBeNil)
// 			So(len(results), ShouldEqual, 0)

// 			results = table.Get(func(order Order) bool {
// 				return order.Account == "1001"
// 			})
// 			So(len(results), ShouldEqual, 0)
// 			results, err = table.GetWithIdx("Account", "1001")
// 			So(err, ShouldNotBeNil)
// 			So(len(results), ShouldEqual, 0)

// 			results = table.Get(func(order Order) bool {
// 				return order.StockCode == "700"
// 			})
// 			So(len(results), ShouldEqual, 0)
// 			results, err = table.GetWithIdx("StockCode", "700")
// 			So(err, ShouldNotBeNil)
// 			So(len(results), ShouldEqual, 0)
// 		})
// 		Convey("Have Index GetWithIdxIntersection", func() {
// 			type Order struct {
// 				ID        string `idx:"normal"`
// 				Account   string `idx:"normal"`
// 				StockCode string `idx:"normal"`
// 				Currency  string `idx:"normal"`
// 				Amount    float64
// 			}

// 			table := New[Order]()
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1002",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1002",
// 				StockCode: "800",
// 				Currency:  "USD",
// 				Amount:    55000,
// 			})

// 			list, err := table.GetWithIdxIntersection(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "700",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 1)

// 			list, err = table.GetWithIdxIntersection(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 2)

// 			list, err = table.GetWithIdxIntersection(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 				"Currency":  "SGD",
// 			})
// 			So(err, ShouldNotBeNil)
// 			So(len(list), ShouldEqual, 0)

// 			list, err = table.GetWithIdxIntersection(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 				"Currency":  "USD",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 0)

// 			list, err = table.GetWithIdxIntersection(map[string]interface{}{
// 				"Account":   "1002",
// 				"StockCode": "800",
// 				"Currency":  "USD",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 1)
// 		})
// 		Convey("Have Index GetWithIdxAggregate", func() {
// 			type Order struct {
// 				ID        string `idx:"normal"`
// 				Account   string `idx:"normal"`
// 				StockCode string `idx:"normal"`
// 				Currency  string `idx:"normal"`
// 				Amount    float64
// 			}

// 			table := New[Order]()
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1001",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1002",
// 				StockCode: "800",
// 				Currency:  "HKD",
// 				Amount:    55000,
// 			})
// 			table.Add(Order{
// 				ID:        "order_1",
// 				Account:   "1002",
// 				StockCode: "800",
// 				Currency:  "USD",
// 				Amount:    55000,
// 			})

// 			list, err := table.GetWithIdxAggregate(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "700",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 3)

// 			list, err = table.GetWithIdxAggregate(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 5)

// 			list, err = table.GetWithIdxAggregate(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 				"Currency":  "SGD",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 5)

// 			list, err = table.GetWithIdxAggregate(map[string]interface{}{
// 				"Account":   "1001",
// 				"StockCode": "800",
// 				"Currency":  "USD",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 5)

// 			list, err = table.GetWithIdxAggregate(map[string]interface{}{
// 				"Account":   "1002",
// 				"StockCode": "800",
// 				"Currency":  "USD",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 4)

// 			type student struct {
// 				Name  string `idx:"normal"`
// 				Age   int    `idx:"normal"`
// 				Class string `idx:"normal"`
// 			}

// 			t := New[student]()
// 			t.Add(student{
// 				Name:  "A",
// 				Age:   18,
// 				Class: "1",
// 			})
// 			t.Add(student{
// 				Name:  `b`,
// 				Age:   18,
// 				Class: "1",
// 			})
// 			t.Add(student{
// 				Name:  `c`,
// 				Age:   20,
// 				Class: "2",
// 			})
// 			t.Add(student{
// 				Name:  `d`,
// 				Age:   17,
// 				Class: "3",
// 			})
// 			list2, err := t.GetWithIdxAggregate(map[string]interface{}{
// 				"Name": "a",
// 				"Age":  18,
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 2)

// 			list2, err = t.GetWithIdxAggregate(map[string]interface{}{
// 				"Class": "1",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 2)

// 			list2, err = t.GetWithIdxAggregate(map[string]interface{}{
// 				"Name": "d",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 1)

// 			list2, err = t.GetWithIdxAggregate(map[string]interface{}{
// 				"Age":   17,
// 				"Class": "0",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 1)

// 			list2, err = t.GetWithIdxAggregate(map[string]interface{}{
// 				"Age":   17,
// 				"Class": "1",
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 3)

// 			list2, err = t.GetWithIdxIntersection(map[string]interface{}{
// 				"Name": "d",
// 				"Age":  20,
// 			})
// 			So(err, ShouldBeNil)
// 			So(len(list2), ShouldEqual, 0)
// 		})
// 	})
// }

// func Test_Clear(t *testing.T) {
// 	Convey("Clear", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()
// 		for i := 0; i < 10000; i++ {
// 			ultraTable.Add(Order{
// 				ID:        fmt.Sprint(i),
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    500.1,
// 			})
// 		}
// 		So(ultraTable.Len(), ShouldEqual, 10000)
// 		So(ultraTable.Cap(), ShouldEqual, 10000)
// 		ultraTable.Clear()
// 		So(ultraTable.Len(), ShouldEqual, 0)
// 		So(ultraTable.Cap(), ShouldEqual, 0)

// 		for i := 0; i < 10000; i++ {
// 			ultraTable.Add(Order{
// 				ID:        fmt.Sprint(i),
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    500.1,
// 			})
// 		}
// 		So(ultraTable.Len(), ShouldEqual, 10000)
// 		So(ultraTable.Cap(), ShouldEqual, 10000)
// 	})
// }

// func Test_Remove(t *testing.T) {
// 	Convey("Remove", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 			At        time.Time
// 		}
// 		Convey("Remove-1", func() {
// 			ultraTable := New[Order]()
// 			ultraTable.Add(Order{
// 				ID:        "1",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    100,
// 			})
// 			ultraTable.Add(Order{
// 				ID:        "1",
// 				Account:   "1001",
// 				StockCode: "9988",
// 				Currency:  "HKD",
// 				Amount:    100,
// 			})
// 			So(ultraTable.RemoveWithIdx(`ID`, "1"), ShouldEqual, 2)
// 			So(ultraTable.Len(), ShouldEqual, 0)
// 			So(ultraTable.RemoveWithIdx(`ID`, "1"), ShouldEqual, 0)
// 			So(ultraTable.RemoveWithIdx(`Account`, `1001`), ShouldEqual, 0)
// 			So(ultraTable.RemoveWithIdx(`StockCode`, `700`), ShouldEqual, 0)
// 			So(ultraTable.Len(), ShouldEqual, 0)

// 			ultraTable.Add(Order{
// 				ID:        "2",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    100,
// 			})
// 			So(ultraTable.Len(), ShouldEqual, 1)
// 			ultraTable.Add(Order{
// 				ID:        "3",
// 				Account:   "1001",
// 				StockCode: "700",
// 				Currency:  "HKD",
// 				Amount:    100,
// 			})
// 			So(ultraTable.Len(), ShouldEqual, 2)
// 			So(ultraTable.RemoveWithIdx(`Account`, `1001`), ShouldEqual, 2)
// 			So(ultraTable.RemoveWithIdx(`StockCode`, `700`), ShouldEqual, 0)
// 		})

// 		Convey("Remove-2", func() {
// 			ultraTable := New[Order]()
// 			for i := 0; i < 1000; i++ {
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: "700",
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 1000)
// 			for i := 0; i < 1000; i++ {
// 				items := ultraTable.Get(func(item Order) bool {
// 					return item.ID == fmt.Sprint(i)
// 				})
// 				So(items[0].Amount, ShouldEqual, i)

// 				isFound := ultraTable.Has(func(item Order) bool {
// 					return item.ID == fmt.Sprint(i)
// 				})
// 				So(isFound, ShouldBeTrue)
// 			}
// 			for i := 0; i < 1000; i++ {
// 				items, err := ultraTable.GetWithIdx("ID", fmt.Sprint(i))
// 				So(err, ShouldBeNil)
// 				So(items[0].Amount, ShouldEqual, i)

// 				isFound := ultraTable.HasWithIdx("ID", fmt.Sprint(i))
// 				So(isFound, ShouldBeTrue)
// 			}

// 			items := ultraTable.GetAll()
// 			So(len(items), ShouldEqual, 1000)

// 			for i := 0; i < 500; i++ {
// 				count := ultraTable.Remove(func(item Order) bool {
// 					return item.ID == fmt.Sprint(i)
// 				})
// 				So(count, ShouldEqual, 1)
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 500)
// 			So(ultraTable.Cap(), ShouldEqual, 1000)
// 			for i := 500; i < 1000; i++ {
// 				count := ultraTable.RemoveWithIdx("ID", fmt.Sprint(i))
// 				So(count, ShouldEqual, 1)
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 0)
// 			So(ultraTable.Cap(), ShouldEqual, 1000)

// 			items = ultraTable.GetAll()
// 			So(len(items), ShouldEqual, 0)
// 		})
// 		Convey("Remove-3", func() {
// 			ultraTable := New[Order]()

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}

// 			for i := 0; i < 10; i++ {
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: "00001",
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 1010)
// 			So(ultraTable.Cap(), ShouldEqual, 1010)

// 			list, err := ultraTable.GetWithIdx("StockCode", "00001")
// 			So(err, ShouldBeNil)
// 			So(len(list), ShouldEqual, 10)

// 			count := ultraTable.RemoveWithIdx("StockCode", "00001")
// 			So(count, ShouldEqual, 10)
// 			So(ultraTable.Len(), ShouldEqual, 1000)
// 			So(ultraTable.Cap(), ShouldEqual, 1010)

// 			for i := 0; i < 10; i++ {
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: "00001",
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 1010)
// 			So(ultraTable.Cap(), ShouldEqual, 1010)

// 			for i := 0; i < 10; i++ {
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: "00001",
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 1020)
// 			So(ultraTable.Cap(), ShouldEqual, 1020)
// 		})
// 		Convey("Remove-4", func() {
// 			ultraTable := New[Order]()

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 					At:        time.Now().Add(-time.Hour),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 500)
// 			So(ultraTable.Cap(), ShouldEqual, 500)
// 			ultraTable.Remove(func(i Order) bool {
// 				return i.At.Before(time.Now())
// 			})
// 			So(ultraTable.Len(), ShouldEqual, 0)
// 			So(ultraTable.Cap(), ShouldEqual, 500)

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 					At:        time.Now().Add(-time.Hour),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 500)
// 			So(ultraTable.Cap(), ShouldEqual, 500)
// 		})
// 		Convey("Remove-5", func() {
// 			ultraTable := New[Order]()

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 					At:        time.Now().Add(-time.Hour),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 500)
// 			So(ultraTable.Cap(), ShouldEqual, 500)
// 			ultraTable.Remove(func(i Order) bool {
// 				return i.At.Before(time.Now())
// 			})
// 			So(ultraTable.Len(), ShouldEqual, 0)
// 			So(ultraTable.Cap(), ShouldEqual, 500)

// 			ultraTable.Clear()

// 			So(ultraTable.Len(), ShouldEqual, 0)
// 			So(ultraTable.Cap(), ShouldEqual, 0)

// 			for i := 0; i < 500; i++ {
// 				rand.Seed(time.Now().UnixNano())
// 				ultraTable.Add(Order{
// 					ID:        fmt.Sprint(i),
// 					Account:   "1001",
// 					StockCode: fmt.Sprint(rand.Intn(1000)),
// 					Currency:  "HKD",
// 					Amount:    float64(i),
// 					At:        time.Now().Add(-time.Hour),
// 				})
// 			}
// 			So(ultraTable.Len(), ShouldEqual, 500)
// 			So(ultraTable.Cap(), ShouldEqual, 500)

// 			for i := 0; i < 500; i++ {
// 				dests, err := ultraTable.GetWithIdx("ID", fmt.Sprint(i))
// 				So(err, ShouldBeNil)
// 				So(dests[0].ID, ShouldEqual, fmt.Sprint(i))
// 			}
// 		})
// 	})
// }

// func Test_Update(t *testing.T) {
// 	Convey("Update", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()
// 		ultraTable.Add(Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})

// 		orders, err := ultraTable.GetWithIdx("ID", "order_1")
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.1)

// 		count := ultraTable.UpdateWithIdx("ID", "order_1", Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})
// 		So(count, ShouldEqual, 1)

// 		orders, err = ultraTable.GetWithIdx("ID", "order_1")
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.2)

// 		ultraTable.Add(Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})

// 		count = ultraTable.UpdateWithIdx("ID", "order_2", Order{
// 			ID:        "order_3",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})
// 		So(count, ShouldEqual, 3)

// 		orders, err = ultraTable.GetWithIdx("ID", "order_3")
// 		So(err, ShouldBeNil)
// 		So(len(orders), ShouldEqual, 3)
// 		So(orders[0].Amount, ShouldEqual, 500.2)
// 		So(orders[1].Amount, ShouldEqual, 500.2)
// 		So(orders[2].Amount, ShouldEqual, 500.2)

// 		count = ultraTable.UpdateWithIdx("StockCode", "700", Order{
// 			ID:        "order_3",
// 			Account:   "1001",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})
// 		So(count, ShouldEqual, 4)

// 	})
// }

// func Test_Type(t *testing.T) {
// 	Convey("type", t, func() {

// 		Convey(`case-1`, func() {
// 			type TypeStruct struct {
// 				A string     `idx:"normal"`
// 				B int        `idx:"normal"`
// 				C int8       `idx:"normal"`
// 				D int16      `idx:"normal"`
// 				E int32      `idx:"normal"`
// 				F int64      `idx:"normal"`
// 				G uint       `idx:"normal"`
// 				H uint8      `idx:"normal"`
// 				I uint16     `idx:"normal"`
// 				J uint32     `idx:"normal"`
// 				K uint64     `idx:"normal"`
// 				L float32    `idx:"normal"`
// 				M float64    `idx:"normal"`
// 				N complex64  `idx:"normal"`
// 				O complex128 `idx:"normal"`
// 				P byte       `idx:"normal"`
// 				Q rune       `idx:"normal"`
// 			}
// 			ultraTable := New[TypeStruct]()
// 			for i := 0; i < 10; i++ {
// 				err := ultraTable.Add(TypeStruct{
// 					A: `test`,
// 					B: int(i),
// 					C: int8(i),
// 					D: int16(i),
// 					E: int32(i),
// 					F: int64(i),
// 					G: uint(i),
// 					H: uint8(i),
// 					I: uint16(i),
// 					J: uint32(i),
// 					K: uint64(i),
// 					L: float32(i),
// 					M: float64(i),
// 					N: 0,
// 					O: 0,
// 					P: byte(i),
// 					Q: rune(i),
// 				})
// 				So(err, ShouldBeNil)
// 			}
// 		})

// 		Convey(`case-2`, func() {
// 			type TypeStruct struct {
// 				A string      `idx:"normal"`
// 				B int         `idx:"normal"`
// 				C int8        `idx:"normal"`
// 				D int16       `idx:"normal"`
// 				E int32       `idx:"normal"`
// 				F int64       `idx:"normal"`
// 				G uint        `idx:"normal"`
// 				H uint8       `idx:"normal"`
// 				I uint16      `idx:"normal"`
// 				J uint32      `idx:"normal"`
// 				K uint64      `idx:"normal"`
// 				L float32     `idx:"normal"`
// 				M float64     `idx:"normal"`
// 				N complex64   `idx:"normal"`
// 				O complex128  `idx:"normal"`
// 				P byte        `idx:"normal"`
// 				Q rune        `idx:"normal"`
// 				X interface{} `idx:"normal"`
// 			}
// 			ultraTable := New[TypeStruct]()
// 			for i := 0; i < 10; i++ {
// 				err := ultraTable.Add(TypeStruct{
// 					A: `test`,
// 					B: int(i),
// 					C: int8(i),
// 					D: int16(i),
// 					E: int32(i),
// 					F: int64(i),
// 					G: uint(i),
// 					H: uint8(i),
// 					I: uint16(i),
// 					J: uint32(i),
// 					K: uint64(i),
// 					L: float32(i),
// 					M: float64(i),
// 					N: 0,
// 					O: 0,
// 					P: byte(i),
// 					Q: rune(i),
// 				})
// 				So(err, ShouldNotBeNil)
// 			}
// 		})

// 		Convey(`case-3`, func() {
// 			type TypeStruct struct {
// 				A string     `idx:"normal"`
// 				B int        `idx:"normal"`
// 				C int8       `idx:"normal"`
// 				D int16      `idx:"normal"`
// 				E int32      `idx:"normal"`
// 				F int64      `idx:"normal"`
// 				G uint       `idx:"normal"`
// 				H uint8      `idx:"normal"`
// 				I uint16     `idx:"normal"`
// 				J uint32     `idx:"normal"`
// 				K uint64     `idx:"normal"`
// 				L float32    `idx:"normal"`
// 				M float64    `idx:"normal"`
// 				N complex64  `idx:"normal"`
// 				O complex128 `idx:"normal"`
// 				P byte       `idx:"normal"`
// 				Q rune       `idx:"normal"`
// 				X []string   `idx:"normal"`
// 			}
// 			ultraTable := New[TypeStruct]()
// 			for i := 0; i < 10; i++ {
// 				err := ultraTable.Add(TypeStruct{
// 					A: `test`,
// 					B: int(i),
// 					C: int8(i),
// 					D: int16(i),
// 					E: int32(i),
// 					F: int64(i),
// 					G: uint(i),
// 					H: uint8(i),
// 					I: uint16(i),
// 					J: uint32(i),
// 					K: uint64(i),
// 					L: float32(i),
// 					M: float64(i),
// 					N: 0,
// 					O: 0,
// 					P: byte(i),
// 					Q: rune(i),
// 				})
// 				So(err, ShouldNotBeNil)
// 			}
// 		})

// 	})
// }

// func Test_SaveWithIdx(t *testing.T) {
// 	Convey("SaveWithIdx", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()

// 		count := ultraTable.SaveWithIdx(`ID`, "order_1", Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})

// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		orders, err := ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.1)

// 		count = ultraTable.SaveWithIdx(`ID`, "order_1", Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})

// 		So(count, ShouldEqual, 1)

// 		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.2)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.SaveWithIdx(`Account`, `1001`, Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.3,
// 		})

// 		So(count, ShouldEqual, 1)

// 		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.3)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.SaveWithIdx(`Account`, `1002`, Order{
// 			ID:        "order_1",
// 			Account:   "1002",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.5,
// 		})

// 		So(count, ShouldEqual, 1)

// 		orders, err = ultraTable.GetWithIdx(`Account`, `1002`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.5)
// 		So(ultraTable.Len(), ShouldEqual, 2)
// 	})
// }

// func Test_SaveWithIdxAggregateAndIntersection(t *testing.T) {
// 	Convey("SaveWithIdx", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()

// 		count := ultraTable.SaveWithIdxAggregate(map[string]interface{}{`ID`: "order_1"}, Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})

// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`ID`: "order_2"}, Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})

// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 2)

// 		list, err := ultraTable.GetWithIdxIntersection(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`})
// 		So(err, ShouldBeNil)
// 		So(len(list), ShouldEqual, 1)

// 		list, err = ultraTable.GetWithIdxAggregate(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`})
// 		So(err, ShouldBeNil)
// 		So(len(list), ShouldEqual, 2)

// 		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`}, Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 2)

// 		for idx, v := range ultraTable.GetAll() {
// 			if idx == 0 {
// 				So(v.Amount, ShouldEqual, 500.2)
// 			}
// 			if idx == 1 {
// 				So(v.Amount, ShouldEqual, 500.1)
// 			}
// 		}
// 		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`Account`: `1001`, `StockCode`: `700`}, Order{
// 			ID:        "order_2",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.3,
// 		})
// 		So(count, ShouldEqual, 2)
// 		So(ultraTable.Len(), ShouldEqual, 2)

// 		count = ultraTable.SaveWithIdxAggregate(map[string]interface{}{`Account`: `1003`, `StockCode`: `900`}, Order{
// 			ID:        "order_3",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.4,
// 		})

// 		for idx, v := range ultraTable.GetAll() {
// 			if idx == 0 {
// 				So(v.ID, ShouldEqual, "order_2")
// 				So(v.Amount, ShouldEqual, 500.3)
// 			}
// 			if idx == 1 {
// 				So(v.ID, ShouldEqual, "order_2")
// 				So(v.Amount, ShouldEqual, 500.3)
// 			}
// 			if idx == 2 {
// 				So(v.ID, ShouldEqual, "order_3")
// 				So(v.Amount, ShouldEqual, 500.4)
// 			}
// 		}

// 		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1005`, `StockCode`: `900`}, Order{
// 			ID:        "order_4",
// 			Account:   "1005",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.5,
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 4)

// 		for idx, v := range ultraTable.GetAll() {
// 			if idx == 0 {
// 				So(v.ID, ShouldEqual, "order_2")
// 				So(v.Amount, ShouldEqual, 500.3)
// 			}
// 			if idx == 1 {
// 				So(v.ID, ShouldEqual, "order_2")
// 				So(v.Amount, ShouldEqual, 500.3)
// 			}
// 			if idx == 2 {
// 				So(v.ID, ShouldEqual, "order_3")
// 				So(v.Amount, ShouldEqual, 500.4)
// 			}
// 			if idx == 3 {
// 				So(v.ID, ShouldEqual, "order_4")
// 				So(v.Amount, ShouldEqual, 500.5)
// 			}
// 		}

// 		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1005`, `StockCode`: `1000`}, Order{
// 			ID:        "order_5",
// 			Account:   "1005",
// 			StockCode: "1000",
// 			Currency:  "HKD",
// 			Amount:    500.5,
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 5)

// 		count = ultraTable.SaveWithIdxIntersection(map[string]interface{}{`Account`: `1005`, `StockCode`: `800`}, Order{
// 			ID:        "order_6",
// 			Account:   "1006",
// 			StockCode: "1100",
// 			Currency:  "HKD",
// 			Amount:    500.6,
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 6)

// 		type StaticQuote struct {
// 			AccountChannel string `idx:"normal"`
// 			CounterID      string `idx:"normal"`
// 			Plevel         string
// 			ImFactor       float64
// 			MmFactor       float64
// 			FmFactor       float64
// 		}
// 		ultraTable2 := New[StaticQuote]()
// 		for i := 0; i < 100; i++ {
// 			counterID := fmt.Sprintf(`ST/HK/%v`, i)
// 			list, _ := ultraTable2.GetWithIdxIntersection(map[string]interface{}{`AccountChannel`: "pspl_sg", `CounterID`: counterID})
// 			So(len(list), ShouldEqual, 0)

// 			ultraTable2.SaveWithIdxIntersection(map[string]interface{}{`AccountChannel`: "pspl_sg", `CounterID`: counterID}, StaticQuote{
// 				AccountChannel: "pspl_sg",
// 				CounterID:      counterID,
// 				Plevel:         "A",
// 				ImFactor:       0.5,
// 				MmFactor:       0.4,
// 				FmFactor:       0.3,
// 			})

// 			So(ultraTable2.Len(), ShouldEqual, i+1)

// 			list, _ = ultraTable2.GetWithIdxIntersection(map[string]interface{}{`AccountChannel`: "pspl_sg", `CounterID`: counterID})
// 			So(len(list), ShouldEqual, 1)
// 		}
// 		So(ultraTable2.Len(), ShouldEqual, 100)

// 		for i := 0; i < 100; i++ {
// 			counterID := fmt.Sprintf(`ST/HK/%v`, i)
// 			ultraTable2.SaveWithIdxIntersection(map[string]interface{}{`AccountChannel`: `lb`, `CounterID`: counterID}, StaticQuote{
// 				AccountChannel: `lb`,
// 				CounterID:      counterID,
// 				Plevel:         "A",
// 				ImFactor:       0.5,
// 				MmFactor:       0.4,
// 				FmFactor:       0.3,
// 			})
// 			So(ultraTable2.Len(), ShouldEqual, 100+i+1)

// 			list, _ := ultraTable2.GetWithIdxIntersection(map[string]interface{}{`AccountChannel`: `lb`, `CounterID`: counterID})
// 			So(len(list), ShouldEqual, 1)
// 		}
// 		So(ultraTable2.Len(), ShouldEqual, 200)
// 	})
// }

// func Test_Kind(t *testing.T) {
// 	Convey("struct", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()
// 		ultraTable.Add(Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})
// 		orders, err := ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.1)

// 		order := orders[0]
// 		order.Amount = 500.2

// 		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.1)

// 		count := ultraTable.UpdateWithIdx("Account", "1001", Order{
// 			ID:        "order_3",
// 			Account:   "1001",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    500.2,
// 		})
// 		So(count, ShouldEqual, 1)

// 		orders, err = ultraTable.GetWithIdx(`Account`, `1001`)
// 		So(err, ShouldBeNil)
// 		So(orders[0].Amount, ShouldEqual, 500.2)
// 	})

// 	Convey("ptr", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[*Order]()
// 		err := ultraTable.Add(&Order{
// 			ID:        "order_1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500.1,
// 		})
// 		So(err, ShouldEqual, OnlySupportPtr)

// 	})
// }

// func Test_HasWithIdx(t *testing.T) {
// 	type Order struct {
// 		ID        string `idx:"normal"`
// 		Account   string `idx:"normal"`
// 		StockCode string `idx:"normal"`
// 		Currency  string
// 		Amount    float64
// 	}
// 	ultraTable := New[Order]()

// 	ultraTable.Add(Order{
// 		ID:        "order_1",
// 		Account:   "1001",
// 		StockCode: "700",
// 		Currency:  "HKD",
// 		Amount:    500.1,
// 	})

// 	ultraTable.Add(Order{
// 		ID:        "order_2",
// 		Account:   "1001",
// 		StockCode: "700",
// 		Currency:  "HKD",
// 		Amount:    500.1,
// 	})

// 	ultraTable.Add(Order{
// 		ID:        "order_3",
// 		Account:   "1002",
// 		StockCode: "700",
// 		Currency:  "HKD",
// 		Amount:    500.1,
// 	})

// 	ultraTable.Add(Order{
// 		ID:        "order_4",
// 		Account:   "1002",
// 		StockCode: "700",
// 		Currency:  "HKD",
// 		Amount:    500.1,
// 	})

// 	ultraTable.Add(Order{
// 		ID:        "order_5",
// 		Account:   "1002",
// 		StockCode: "700",
// 		Currency:  "HKD",
// 		Amount:    500.1,
// 	})

// 	Convey("HasWithIdx", t, func() {
// 		So(ultraTable.HasWithIdx(`ID`, "order_1"), ShouldBeTrue)
// 		So(ultraTable.HasWithIdx(`Account`, `1002`), ShouldBeTrue)
// 		So(ultraTable.HasWithIdx(`Account`, `1003`), ShouldBeFalse)
// 		So(ultraTable.HasWithIdx(`Currency`, `HKD`), ShouldBeFalse)
// 	})
// 	Convey("GetWithIdxCount", t, func() {
// 		So(ultraTable.GetWithIdxCount(`ID`, "order_1"), ShouldEqual, 1)
// 		So(ultraTable.GetWithIdxCount(`Account`, `1002`), ShouldEqual, 3)
// 		So(ultraTable.GetWithIdxCount(`Account`, `1003`), ShouldEqual, 0)
// 		So(ultraTable.GetWithIdxCount(`Currency`, `HKD`), ShouldEqual, 0)
// 	})
// 	Convey("GetWithIdxAggregateCount", t, func() {
// 		So(ultraTable.GetWithIdxAggregateCount(map[string]interface{}{
// 			`ID`:      "order_1",
// 			`Account`: `1001`,
// 		}), ShouldEqual, 2)
// 		So(ultraTable.GetWithIdxAggregateCount(map[string]interface{}{
// 			`ID`:      "order_2",
// 			`Account`: `1001`,
// 		}), ShouldEqual, 2)
// 		So(ultraTable.GetWithIdxAggregateCount(map[string]interface{}{
// 			`ID`:      "order_6",
// 			`Account`: `1001`,
// 		}), ShouldEqual, 2)

// 		So(ultraTable.GetWithIdxAggregateCount(map[string]interface{}{
// 			`ID`:      "order_6",
// 			`Account`: `1003`,
// 		}), ShouldEqual, 0)
// 	})
// 	Convey("GetWithIdxIntersectionCount", t, func() {
// 		So(ultraTable.GetWithIdxIntersectionCount(map[string]interface{}{
// 			`ID`:      "order_1",
// 			`Account`: `1001`,
// 		}), ShouldEqual, 1)
// 		So(ultraTable.GetWithIdxIntersectionCount(map[string]interface{}{
// 			`ID`:      "order_2",
// 			`Account`: `1002`,
// 		}), ShouldEqual, 0)
// 		So(ultraTable.GetWithIdxIntersectionCount(map[string]interface{}{
// 			`ID`:      "order_6",
// 			`Account`: `1001`,
// 		}), ShouldEqual, 0)

// 		So(ultraTable.GetWithIdxIntersectionCount(map[string]interface{}{
// 			`ID`:      "order_6",
// 			`Account`: `1003`,
// 		}), ShouldEqual, 0)
// 	})

// 	Convey("revome", t, func() {
// 		count := ultraTable.RemoveWithIdx("ID", "order_1")
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.HasWithIdx("ID", "order_1"), ShouldBeFalse)
// 	})
// }

// func Test_Copy(t *testing.T) {
// 	Convey("copy", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}

// 		ultraTable := New[Order]()
// 		order := Order{
// 			ID:        "1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    100.1,
// 		}

// 		ultraTable.Add(order)
// 		fmt.Printf("%p \r\n", &order)

// 		dest, _ := ultraTable.GetWithIdx("ID", "1")

// 		orderBefore := dest[0]
// 		orderBefore.Amount = 100.2

// 		fmt.Printf("%p \r\n", &orderBefore)

// 		dest, _ = ultraTable.GetWithIdx("ID", "1")
// 		So(dest[0].Amount, ShouldEqual, 100.1)

// 		fmt.Printf("%p \r\n", &dest[0])

// 	})
// }

// func Test_RemoveWithIdxIntersection(t *testing.T) {
// 	Convey("RemoveWithIdxIntersection", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()

// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1002",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1002",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})

// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1003",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1004",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})

// 		count := ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
// 			"Account":   "1001",
// 			"StockCode": "800",
// 		})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 5)

// 		count = ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
// 			"Account":   "1002",
// 			"StockCode": "700",
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 4)

// 		count = ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
// 			"StockCode": "700",
// 		})
// 		So(count, ShouldEqual, 2)
// 		So(ultraTable.Len(), ShouldEqual, 2)

// 		count = ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
// 			"Account": "1003",
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
// 			"Account":   "1002",
// 			"ID":        "1",
// 			"StockCode": "800",
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 0)

// 		count = ultraTable.RemoveWithIdxIntersection(map[string]interface{}{})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 0)
// 	})
// }

// func Test_RemoveWithIdxAggregate(t *testing.T) {
// 	Convey("RemoveWithIdxAggregate", t, func() {
// 		type Order struct {
// 			ID        string `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}
// 		ultraTable := New[Order]()

// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1001",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1002",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1002",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})

// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1003",
// 			StockCode: "800",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})
// 		ultraTable.Add(Order{
// 			ID:        "1",
// 			Account:   "1004",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    55000,
// 		})

// 		count := ultraTable.RemoveWithIdxAggregate(map[string]interface{}{
// 			"Account":   "1005",
// 			"StockCode": "600",
// 		})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 5)

// 		count = ultraTable.RemoveWithIdxAggregate(map[string]interface{}{
// 			"Account":   "1002",
// 			"StockCode": "700",
// 		})
// 		So(count, ShouldEqual, 4)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.RemoveWithIdxAggregate(map[string]interface{}{
// 			"StockCode": "700",
// 		})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 1)

// 		count = ultraTable.RemoveWithIdxAggregate(map[string]interface{}{
// 			"Account": "1003",
// 		})
// 		So(count, ShouldEqual, 1)
// 		So(ultraTable.Len(), ShouldEqual, 0)

// 		count = ultraTable.RemoveWithIdxAggregate(map[string]interface{}{
// 			"Account":   "1002",
// 			"ID":        "1",
// 			"StockCode": "800",
// 		})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 0)

// 		count = ultraTable.RemoveWithIdxAggregate(map[string]interface{}{})
// 		So(count, ShouldEqual, 0)
// 		So(ultraTable.Len(), ShouldEqual, 0)
// 	})
// }

// func Test_Concurrent(t *testing.T) {
// 	Convey("Concurrent", t, func() {
// 		type Order struct {
// 			ID        int    `idx:"normal"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string
// 			Amount    float64
// 		}

// 		Convey("Concurrent-1", func() {

// 			waitGroup := sync.WaitGroup{}
// 			len := 10
// 			waitGroup.Add(len * 3)

// 			ultraTable := New[Order]()
// 			for i := 0; i < len; i++ {
// 				go func() {
// 					ultraTable.Add(Order{
// 						ID:        i,
// 						Account:   `a`,
// 						StockCode: `700`,
// 						Currency:  `HKD`,
// 						Amount:    100,
// 					})
// 					waitGroup.Done()
// 				}()
// 			}
// 			for i := 0; i < len; i++ {
// 				go func() {
// 					ultraTable.UpdateWithIdx(`ID`, i, Order{
// 						ID:        i,
// 						Account:   `a1`,
// 						StockCode: `800`,
// 						Currency:  `USD`,
// 						Amount:    100,
// 					})
// 					waitGroup.Done()
// 				}()
// 			}

// 			for i := 0; i < len; i++ {
// 				go func() {
// 					ultraTable.GetWithIdx("ID", i)
// 					waitGroup.Done()
// 				}()
// 			}
// 			waitGroup.Wait()
// 			So(ultraTable.Len(), ShouldEqual, len)
// 		})

// 		Convey("Concurrent-Read-Write", func() {
// 			waitGroup := sync.WaitGroup{}
// 			len := 100
// 			waitGroup.Add(len * 2)

// 			ultraTable := New[Order]()
// 			for i := 0; i < len; i++ {
// 				go func() {
// 					ultraTable.Add(Order{
// 						ID:        i,
// 						Account:   `a`,
// 						StockCode: `700`,
// 						Currency:  `HKD`,
// 						Amount:    100,
// 					})
// 					waitGroup.Done()
// 				}()
// 			}

// 			for i := 0; i < len; i++ {
// 				go func() {
// 					ultraTable.GetWithIdx("ID", i)
// 					waitGroup.Done()
// 				}()
// 			}
// 			waitGroup.Wait()
// 			So(ultraTable.Len(), ShouldEqual, len)
// 		})
// 	})
// }

// func Test_UqIndex(t *testing.T) {
// 	Convey("unique index duplicate", t, func() {
// 		type order struct {
// 			ID        string `idx:"unique"`
// 			Account   string
// 			StockCode string
// 			Currency  string
// 			Amount    float64
// 		}

// 		table := New[order]()

// 		err := table.Add(order{
// 			ID:        "10021103212",
// 			Account:   "1111",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldBeNil)

// 		err = table.Add(order{
// 			ID:        "10021103212",
// 			Account:   "1112",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldEqual, UniqueIndex)

// 		err = table.Add(order{
// 			ID:        "10021103213",
// 			Account:   "1112",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldBeNil)

// 		So(table.Len(), ShouldEqual, 2)
// 		So(table.Cap(), ShouldEqual, 2)
// 	})

// 	Convey("comprehensive index duplicate", t, func() {
// 		type order struct {
// 			ID        string `idx:"unique"`
// 			Account   string `idx:"normal"`
// 			StockCode string `idx:"normal"`
// 			Currency  string `idx:"normal"`
// 			Amount    float64
// 		}

// 		table := New[order]()

// 		err := table.Add(order{
// 			ID:        "10021103212",
// 			Account:   "1111",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldBeNil)

// 		err = table.Add(order{
// 			ID:        "10021103212",
// 			Account:   "1112",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldEqual, UniqueIndex)

// 		err = table.Add(order{
// 			ID:        "10021103213",
// 			Account:   "1112",
// 			StockCode: "700",
// 			Currency:  "HKD",
// 			Amount:    500,
// 		})
// 		So(err, ShouldBeNil)

// 		So(table.Len(), ShouldEqual, 2)
// 		So(table.Cap(), ShouldEqual, 2)
// 	})
// }
