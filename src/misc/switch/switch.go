package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3}
	i := 2

	switch i {
	case x[1]:
		fmt.Println("a")
	case 1, 3:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}

	x1 := 10
	switch x1 {
	case 10:
		fmt.Println("a")
		fallthrough
	case 0:
		fmt.Println("b")
	}

	// 省略条件表达式，可当 if...else if...else 使用
	switch {
	case x[1] > 0:
		fmt.Println("aa")
	case x[1] < 0:
		fmt.Println("bb")
	default:
		fmt.Println("cc")
	}

	switch i := x[2]; {
	// 带初始化语句
	case i > 0:
		fmt.Println("a")
	case i < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
