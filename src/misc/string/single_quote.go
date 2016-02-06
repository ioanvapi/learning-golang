package main

import (
	"fmt"
)

// 单引号字符常量表示 Unicode Code Point，支持 \uFFFF、\U7FFFFFFF、\xFF 格式。对应 rune 类型，UCS-4。
func main() {
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'

	fmt.Println(c1 == '我', string(c2) == "\xe4\xbb\xac")
}
