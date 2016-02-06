package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	fmt.Println(a, b, c, d)

	// 内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。
	x1 := []int{0, 0, 0} // 提供初始化表达式
	x1[1] = 10
	fmt.Println(x1)

	x2 := make([]int, 3) // make slice
	x2[1] = 10
	fmt.Println(x2)

	x3 := new([]int)
	// x3[0] = 10 // Error: invalid operation: c[1] (index of type *[]int)
	fmt.Println(x3)
}
