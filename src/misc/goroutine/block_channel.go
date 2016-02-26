package main

import (
	"fmt"
)

var ch = make(chan int)

func foo() {
	fmt.Println("foo: before chan")
	ch <- 0 // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
	fmt.Println("foo: after chan")
}

func main() {
	go foo()
	fmt.Println("main: before main")
	<-ch // 从ch取数据，如果ch中还没放数据，那就挂起main线，直到foo函数中放数据为止
	fmt.Println("main: after main")
}
