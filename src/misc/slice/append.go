package main

import (
	"fmt"
)

// 向 slice 尾部添加数据，返回新的 slice 对象

func main() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)

	s1 := append(s, 1)
	fmt.Printf("%p\n", &s1)

	fmt.Println(s, s1)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := data[:3]
	s3 := append(s2, 100, 200) // 添加多个值

	fmt.Println(data)
	fmt.Println(s2)
	fmt.Println(s3)

	// 一旦超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	data1 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s4 := data1[:2:3]

	s4 = append(s4, 100, 200) // 一次 append 两个值，超出 s4.cap 限制。

	fmt.Println(s4, data1)         // 重新分配底层数组，与原数组无关
	fmt.Println(&s4[0], &data1[0]) // 比对底层数组起始指针

	// 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，
	// 以减少内存分配和数据复制开销。或初始化足够长的 len 属性，改用索引号进行操作。
	// 及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
	s5 := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s5 = append(s5, i)
		if n := cap(s5); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
