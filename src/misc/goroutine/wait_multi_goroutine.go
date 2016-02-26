package main

import (
	"fmt"
)

var quit chan int // 只开一个信道

func foo(id int) {
	fmt.Println(id)
	quit <- 0 // ok, finished
}

func main() {
	count := 1000
	// quit = make(chan int)        // 无缓冲
	quit = make(chan int, count) // 容量1000

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		<-quit
	}
}
