package main

import (
	"fmt"
)

// the other init function in this go source file
func init() {
	fmt.Println("do in init")
}

func main() {
	fmt.Println("do in main")
}

func testf() {
	fmt.Println("do in testf")
	//if uncomment the next statment, then go build give error message : .\gprog.go:19: undefined: init
//	init()
}