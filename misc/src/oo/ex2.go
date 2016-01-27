package main

import (
	"fmt"
)

// type declaration (func)
type Foo func() int

// declaring a method
func (f Foo) Add(x int) int {
	return f() + x
}

func main() {
	var x Foo
	x = func() int {return 1}

	fmt.Println(x())
	fmt.Println(x.Add(3))
}