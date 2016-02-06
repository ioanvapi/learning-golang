package main

import (
	"fmt"
)

func main() {
	var num3 = 3.7E+1 + 5.98E-2i

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num3, (num3))
}
