package main

import (
	"fmt"
)

func test (a int) int {
	return (a+2+(a <<1)) %8 ^ 5
}

// 将任何长度字符串, 散列成 0-15 整数.
// 通过 hashcode 散列出的 0-15 的数字的概率是相等的.
func HashCode(key string) int {
	var index int = 0
	index = int(key[0])
	for k :=0 ; k< len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 -1
	return index
}

func main() {
	fmt.Println(test(132223))
	fmt.Println(test(1545323))
	fmt.Println(test(324234))
	fmt.Println(HashCode("sdfsdf"))
}