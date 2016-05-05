package main

import (
	"fmt"
	"time"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		// time.Sleep(1 * time.Second)
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// fmt.Println(f(), f(), f(), f(), f())
	for i := 0; i < 14; i++ {
		fmt.Println(f())
	}
	time.Sleep(10 * time.Second)
}
