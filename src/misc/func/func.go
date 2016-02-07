package main

import (
	"fmt"
)

func test(x, y int, s string) (int, string) {
	// 类型相同的相邻参数可合并
	// 多返回值必须用括号
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func test1(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string // 定义函数类型

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

// 有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
func main() {
	fmt.Println(test(1, 2, "hello"))

	s1 := test1(func() int { return 100 }) // 直接将匿名函数当参数

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println(s1, s2)
}
