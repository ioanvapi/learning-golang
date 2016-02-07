package main

import (
	"fmt"
)

// 类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
func main() {
	s := "abc"

	for i := range s {
		// 忽略 2nd value，支持 string/array/slice/map。
		fmt.Println(s[i])
	}

	for _, c := range s {
		// 忽略 index
		fmt.Println(c)
	}

	for range s {
		// 忽略全部返回值, 仅迭代
		// ...
	}

	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		// 返回 (key, value)
		fmt.Println(k, v)
	}

	a := [3]int{0, 1, 2}

	// 注意 range 会复制对象
	for i, v := range a {
		// index, value 都是从复制品中取出
		if i == 0 {
			// 在修改前, 我们先修改原数组
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效, 输出 [0, 999, 999]
		}

		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组
	}
	fmt.Println(a) // 输出 [100, 101, 102]
}
