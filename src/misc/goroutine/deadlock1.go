package main

import "fmt"

var ch1 = make(chan int)
var ch2 = make(chan int)

func say(s string) {
	fmt.Println(s)
	ch1 <- <-ch2 // ch1 等待 ch2流出的数据
}

func main() {
	go say("hello")
	//ch2 <- 0 // uncomment it cause deadlock
	<-ch1 // 堵塞主线
}
