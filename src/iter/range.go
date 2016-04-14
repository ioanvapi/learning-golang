package main

import (
	"fmt"
	// "github.com/bradfitz/iter"
)

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func main() {
	for i := range N(4) {
		fmt.Println(i)
	}
}
