package main

import (
	"fmt"
)

func main() {
	var x = make([]struct{}, 100)
	var y = x[:50]
	fmt.Println(len(y), cap(y)) // prints 50 100

	var a struct{}
	var b = &a
	fmt.Println(a, b)

	// https://golang.org/ref/spec#Size_and_alignment_guarantees
	var aa, bb struct{}
	fmt.Println(&aa == &bb) // true

	a1 := make([]struct{}, 10)
	b1 := make([]struct{}, 20)
	fmt.Println(&a1 == &b1)       // false, a and b are different slices
	fmt.Println(&a1[0] == &b1[0]) // true, their backing arrays are the same

	a2 := struct{}{} // not the zero value, a real new struct{} instance
	b2 := struct{}{}
	fmt.Println(a2 == b2) // true

}
