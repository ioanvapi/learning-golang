package main

import (
	"fmt"
)

func main() {
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	fmt.Println(string(bs))
	u := "电脑"

	us := []rune(u)
	us[1] = '话'
	fmt.Println(string(us))
}
