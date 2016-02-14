package main

import (
	"fmt"
)

func main() {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10}, // 可省略元素类型
		2: {"user2", 20},
	}
	fmt.Println(m)
	println(m[1].name)

	// 预先给 make 函数一个合理元素数量参数，有助于提升性能。因为事先申请一大块内存，可避免后续操作时频繁扩张。
	m1 := make(map[string]int, 1000)
	fmt.Println(m1)
}
