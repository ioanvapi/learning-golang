package main

import (
	"fmt"
)

func main() {
	s := "abc汉字"
	fmt.Println("len(s): ", len(s))
	for i := 0; i < len(s); i++ {
		// byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s {
		// rune
		fmt.Printf("%c,", r)
	}
	fmt.Println()
}
