package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	slice := data[1:4:5] // [low : high : max]
	fmt.Println(data, slice)
	//创建表达式使用的是元素索引号，而非数量。
	a := []int{0, 1, 2}
	b := [...]int{0, 1, 2}
	fmt.Println(len(a), cap(a), reflect.TypeOf(a))
	fmt.Println(len(b), cap(b), reflect.TypeOf(b))

	data1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := data1[:6:8]
	fmt.Println("slice1", slice1, len(slice1), cap(slice1)) // 省略low

	slice2 := data1[5:]
	fmt.Println("slice2", slice2, len(slice2), cap(slice2)) // 省略high, max

	slice3 := data1[:3]
	fmt.Println("slice3", slice3, len(slice3), cap(slice3)) // 省略low, max

	slice4 := data1[:]
	fmt.Println("slice4", slice4, len(slice4), cap(slice4)) // 全部省略

	// 读写操作实际目标是底层数组，只需注意索引号的差别。
	data2 := [...]int{0, 1, 2, 3, 4, 5}
	s := data2[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data2)

	// 可直接创建 slice 对象，自动分配底层数组。
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))

	// 使用 make 动态创建 slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
	s4 := []int{0, 1, 2, 3}

	p := &s4[2] // *int, 获取底层数组元素指针
	*p += 100

	fmt.Println(s4)

	// 至于 [][]T，是指元素类型为 []T的 slice
	data3 := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data3)

	// 可直接修改 struct array/slice 成员。
	d := [5]struct {
		x int
	}{}

	s5 := d[:]

	d[1].x = 10
	s5[2].x = 20

	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}
