package main

// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

import "fmt"

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func main() {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)

	test(a)
	fmt.Println(a)
}
