package main

import (
	"fmt"
)

func length(s string) int {
	fmt.Println("call length.")
	return len(s)
}

func main() {
	s := "abcd"
	// 不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。
	for i, n := 0, length(s); i < n; i++ {
		// 避免多次调用 length 函数
		fmt.Println(i, s[i])
	}
	// fmt.Println(i, n) // undefined
}
