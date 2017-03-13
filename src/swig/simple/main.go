package main

import (
	"fmt"

	"./simple"
)

func main() {
	// Call our gcd() function
	x := 42
	y := 105
	g := simple.Gcd(x, y)
	fmt.Println("The gcd of", x, "and", y, "is", g)

	// Manipulate the Foo global variable

	// Output its current value
	fmt.Println("Foo =", simple.GetFoo())

	// Change its value
	simple.SetFoo(3.1415926)

	// See if the change took effect
	fmt.Println("Foo =", simple.GetFoo())
}
