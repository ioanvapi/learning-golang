package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("NumCPU: ", runtime.NumCPU())

	r := runtime.GOMAXPROCS(-1)
	fmt.Println("GOMAXPROCS(-1): ", r)

	r = runtime.GOMAXPROCS(2)
	fmt.Println("GOMAXPROCS(2): ", r)

	r = runtime.GOMAXPROCS(8)
	fmt.Println("GOMAXPROCS(8): ", r)
}