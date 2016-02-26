package main

import (
	"fmt"
)

var ch = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println()
}
