package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go say("world") //开一个新的Goroutines执行
	for {
	}
}
