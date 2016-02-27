package main

import (
	"fmt"
)

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Printf("bar: %p, %T\n", p, p)
	fmt.Println(*p)
}

func main() {
	fmt.Printf("main1: %p, %T\n", p, p)
	var err error
	p, err = foo()
	// p, err := foo() // cause nil pointer
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("main2: %p, %T\n", p, p)
	bar()
	fmt.Println(*p)
}
