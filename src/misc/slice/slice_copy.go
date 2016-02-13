package main

import (
	"fmt"
)

func main() {
	// 函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
	// 应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := data[8:]
	s2 := data[:5]

	copy(s2, s) // dst:s2, src:s

	fmt.Println(s2)
	fmt.Println(data)
}
