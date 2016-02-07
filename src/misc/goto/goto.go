package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}
BREAK:
	fmt.Println("break")
	// EXIT: // error
}
