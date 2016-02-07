package main

import (
	"fmt"
)

func main() {
	// func variable
	fn := func() { fmt.Println("Hello, World!") }
	fn()

	// func slice
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}

	fmt.Println(fns[0](100))

	// func as field
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}

	fmt.Println(d.fn())

	// channel of func
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	fmt.Println((<-fc)())
}
