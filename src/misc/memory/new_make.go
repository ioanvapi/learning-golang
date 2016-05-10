package main

import (
	"fmt"
)

func main() {
	var p1 *[]int = new([]int)     // 为切片结构分配内存；*p == nil；很少使用
	var v1 []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用
	fmt.Println(p1)
	fmt.Println(v1)

	// 不必要地使问题复杂化：
	var p *[]int = new([]int)
	fmt.Println(p) //输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p)       //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0

	// 习惯用法:
	v := make([]int, 10)
	fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]
}
