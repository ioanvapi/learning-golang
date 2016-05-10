package main

import (
	"fmt"
)

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("d")
	fmt.Println("e")
}

func main() {
	defer func() {
		// 必须要先声明defer, 否则不能捕获到panic异常
		fmt.Println("b")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容, 55
		}
		fmt.Println("c")
	}()
	f()
}
