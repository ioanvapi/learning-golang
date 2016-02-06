package main

import (
	"fmt"
)

func main() {
	var a struct {
		x int `a`
	}
	var b struct {
		x int `ab`
	}
	// b = a // error
	fmt.Println(a, b)

	type bigint int64

	var x bigint = 100
	fmt.Println(x)
}
