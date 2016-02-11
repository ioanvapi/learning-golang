package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (person *Person) Move(newAddress string) string {
	old := person.Address
	person.Address = newAddress
	return old
}

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}
