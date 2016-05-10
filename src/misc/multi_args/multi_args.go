package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ") // 输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { // 要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
