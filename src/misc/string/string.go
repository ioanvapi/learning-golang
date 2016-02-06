package main

import (
	"fmt"
)

func main() {
	// 声明一个string类型变量并赋值
	//var str1 string = "\\\""
	var str1 = "\\\""

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\"`)

	// 使用索引号访问字符 (byte)
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	s1 := `a
	b\r\n\x00
	c`
	fmt.Println(s1)

	//连接跨行字符串时，"+" 必须在上一行末尾，否则导致编译错误。
	s2 := "Hello, " +
		"World!"
	fmt.Println(s2)

	// s3 := "Hello, "
	// +"World!"

	x := "Hello, World!"
	x1 := x[:5]  // Hello
	x2 := x[7:]  // World!
	x3 := x[1:5] // ello
	fmt.Println(x1, x2, x3)
}
