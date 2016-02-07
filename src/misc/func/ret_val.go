package main

import (
	"fmt"
)

func test() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

func main() {
	// 不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
	// s := make([]int, 2)
	// s = test() // error
	x, _ := test()
	fmt.Println(x)

	fmt.Println(add(test()))
	fmt.Println(sum(test()))
}
