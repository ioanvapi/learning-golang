package main

import "fmt"

var quit = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {

	// runtime.GOMAXPROCS(2) // 最多使用2个核

	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
	fmt.Println()
}
