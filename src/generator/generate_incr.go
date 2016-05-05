package main

import "fmt"

func incr(max int) func() int {
	init := 0
	return func() int {
		if init == max {
			init = 0
		}
		init++
		return init
	}
}

func main() {
	// f := incr(5)
	for i := 0; i < 15; i++ {
		// fmt.Println(f())
		fmt.Println(incr(5)())
	}
}
