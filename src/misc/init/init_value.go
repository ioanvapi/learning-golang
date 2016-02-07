package main

import (
	"fmt"
)

func main() {
	// 初始化复合对象，必须使用类型标签，且左大括号必须在类型尾部。
	// var a struct {x int} = {100}
	// var b []int = {1,2,3}
	// c := struct {x int; y string}
	var a = struct{ x int }{100}
	var b = []int{1, 2, 3}

	fmt.Println(a, b)

	// 初始化值以 "," 分隔。可以分多行，但最后一行必须以 "," 或 "}" 结尾。
	// x := []int {
	// 1,
	// 2 // Error
	// }

	x := []int{
		1,
		2, // ok
	}
	x1 := []int{
		1,
		2} // ok
	fmt.Println(x, x1)
}
