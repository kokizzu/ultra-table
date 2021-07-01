package main

import (
	"fmt"
	"time"

	ultra_table "github.com/longbridgeapp/ultra-table"
)

// func sum(seq int, ch chan int) {
// 	defer close(ch)
// 	sum := 0
// 	for i := 1; i <= 10000000; i++ {
// 		sum += i
// 	}
// 	fmt.Printf("子协程%d运算结果:%d\n", seq, sum)
// 	ch <- sum
// }

// func main() {
// 	// 启动时间
// 	start := time.Now()
// 	// 最大 CPU 核心数
// 	cpus := runtime.NumCPU()
// 	runtime.GOMAXPROCS(cpus)
// 	chs := make([]chan int, cpus)
// 	for i := 0; i < len(chs); i++ {
// 		chs[i] = make(chan int, 1)
// 		go sum(i, chs[i])
// 	}
// 	sum := 0
// 	for _, ch := range chs {
// 		res := <-ch
// 		sum += res
// 	}
// 	// 结束时间
// 	end := time.Now()
// 	// 打印耗时
// 	fmt.Printf("最终运算结果: %d, 执行耗时(s): %f\n", sum, end.Sub(start).Seconds())
// }

type Order struct {
	ID        string `index:"id"`
	Account   string `index:"account"`
	StockCode string `index:"stock_code"`
	Currency  string
	Amount    float64
}

func main() {
	table := ultra_table.NewUltraTable()
	for i := 0; i < 1000000; i++ {
		table.Add(Order{
			ID:        fmt.Sprint(i),
			Account:   `1111`,
			StockCode: `700`,
			Currency:  `HKD`,
			Amount:    10,
		})
	}
	for {
		fmt.Println(`ticker`)
		time.Sleep(time.Second)
		for i := 0; i < 10000; i++ {
			table.Add(Order{
				ID:        fmt.Sprint(i),
				Account:   `1111`,
				StockCode: `700`,
				Currency:  `HKD`,
				Amount:    10,
			})
		}
	}
}
