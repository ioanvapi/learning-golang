package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s string
	var c complex128
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(c))

	var a [3]uint32
	fmt.Println(unsafe.Sizeof(a))

	type SS struct {
		a uint16
		b uint32
	}
	var ss SS
	fmt.Println(unsafe.Sizeof(ss))

	var sss struct{}
	fmt.Println(unsafe.Sizeof(sss))

	type SSSS struct {
		A struct{}
		B struct{}
	}
	var ssss SSSS
	fmt.Println(unsafe.Sizeof(ssss))

	var x [1000000000]struct{}
	fmt.Println(unsafe.Sizeof(x)) // prints 0

	// Slices of struct{}s consume only the space for their slice header
	var xx = make([]struct{}, 1000000000)
	fmt.Println(unsafe.Sizeof(xx))
}
