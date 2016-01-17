package main

import (
	"fmt"
)

func main() {
	s := "Go编程"
	fmt.Println(len(s))                 // 8因为中文字符是用3个字节存的
	fmt.Println(len(string(rune('编')))) // 3
	// 如果想要获得我们想要的情况的话, 需要先转换为rune切片再使用内置的len函数.
	fmt.Println(len([]rune(s))) // 4
	// 所以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。
}
