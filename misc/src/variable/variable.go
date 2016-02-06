package main

import (
	"fmt"
)

var x, y, z int
var s, n = "abc", 123

var (
	a int
	b float32
)

func main() {
	n, s := 0x1234, "Hello, World!"
	fmt.Println(x, s, n)
}
