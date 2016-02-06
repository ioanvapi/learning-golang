package main

import "fmt"

func main() {
	var b byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(b)
	fmt.Println(b)
	fmt.Println(n)
	// a := 100
	// if bool(a) {
	// fmt.Println("true")
	// }
}
