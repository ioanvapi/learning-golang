package main

import (
	"fmt"
)

// 不支持运算符重载, "++", "--"是语句而非表达式
func main() {
	n := 0
	p := &n
	// b := n++
	// if n++ == 1 {}
	// ++n
	n++
	*p++ //(*p)++
	fmt.Println(n, *p)
}
