package main

import "fmt"

var err = 100

func f() {
	fmt.Printf("global err address: %0#x\n", &err)
	err := 101
	fmt.Printf("local err address: %0#x\n", &err)
	fmt.Printf("local err value: %v\n", err)
	a, err := 123, 102
	_ = a
	fmt.Printf("again local err address: %0#x\n", &err)
	fmt.Printf("again local err value: %v\n", err)
}

func main() {
	f()
	fmt.Printf("global err: %v\n", err)
}
