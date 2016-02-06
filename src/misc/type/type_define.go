package main

import (
	"fmt"
)

func main() {
	type bigint int64
	type myslice []int
	x := 1234
	var b bigint = bigint(x) // 命名类型, 必须显示转换, 除非是常量
	var b2 int64 = int64(b)
	fmt.Println(b, b2)

	var s myslice = []int{1, 2, 3} // 未命名类型, 隐式转换
	var s2 []int = s
	fmt.Println(s, s2)
}
