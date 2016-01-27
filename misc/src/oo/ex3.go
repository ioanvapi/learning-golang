package main

import (
	"fmt"
)

type Person struct {
	Name string
	Address
}

type Address struct {
	Number string
	Street string
	City string
	State string
	Zip string
}

// declaring a method
func (a Address) String() string {
	return a.Number + " " + a.Street + "\n" + a.City + ", " + a.State + " " + a.Zip + "\n"
}

// Not overloading
func (p Person) String() string {
	return p.Name + "\n" + p.Address.String()
}

// Types Remain Distinct
func isValidAddress(a Address) bool {
	return a.Street != ""
}

// Declare using Composite Literal
func main() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City: "Gotham",
			State: "NY",
			Zip: "01313",
		},
	}

	fmt.Println(p.String())
	fmt.Println(p.Address.String())

	// cannot use p (type Person) as type Address
	// in argument to isValidAddress
	//fmt.Println(isValidAddress(p))

	fmt.Println(isValidAddress(p.Address))
}