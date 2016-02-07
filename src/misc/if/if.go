package main

import (
	"fmt"
)

func main() {
	x := 0
	// if x > 10
	// {
	// }
	if n := "abc"; x > 10 {
		// 初始化语句未必就是定义变量, 比如, fmt.Println("init")也可以
		fmt.Println(n[2])
	} else if x < 0 {
		// 注意 else if 和 else 左大括号位置
		fmt.Println(n[1])
	} else {
		fmt.Println(n[0])
	}
}
