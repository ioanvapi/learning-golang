package main

import (
	"fmt"
)

// 支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。

func main() {
	type data struct {
		a int
	}
	var d = data{1234}
	var p *data
	p = &d

	fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。

	// x := 1234
	// p1 := &x
	// p++ // error 不能对指针做加减法运算
}
