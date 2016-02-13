package main

import (
	"fmt"
)

func main() {
	// 所谓 reslice，是基于已有 slice 创建新 slice 对象，以便在 cap 允许范围内调整属性。
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 新对象依旧指向原底层数组。
	s1 := s[2:5]
	s1[2] = 100
	fmt.Println(s1)

	s2 := s1[2:6:7]
	s2[3] = 200
	fmt.Println(s2)
	// s3 := s2[3:6]

	fmt.Println(s)
}
