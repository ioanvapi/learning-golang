package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 0x12345678
	p := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p)      // Pointer -> *[4]byte

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	fmt.Println()
}
