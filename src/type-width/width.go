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
}
