package main

import (
	"fmt"
)

// the first init function in this go source file
func init() {
	fmt.Println("do in init1")
}

// the second init function in this go source file
func init() {
	fmt.Println("do in init2")
}

func main() {
	fmt.Println("do in main")
}