package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

//	ch1 <- 1
	select {
	case <- ch1:
		fmt.Println("ch1 pop one element")
	case <- ch2:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}
}