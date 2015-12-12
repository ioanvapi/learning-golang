package main

import (
	"fmt"
)

func main() {
	var (
		num1 int
		num2 int
		num3 int
	)
	num1, num2, num3 = 1, 2, 3
	// 打印函数调用语句。用于打印上述三个变量的值。
	fmt.Println(num1, num2, num3)
}
