package main

import (
	"fmt"
)

func main() {
	s := "abc"

	for i, n := 0, len(s); i < n; i++ {
		// 常见的 for 循环, 支持初始化语句
		fmt.Println(s[i])
	}

	n := len(s)
	for n > 0 {
		// 替代 while (n > 0) {}
		// 替代 for(; n > 0;){}
		n--
		fmt.Println(s[n])
	}

	for {
		// 替代 while(true){}
		// 替代 for (;;) {}
		fmt.Println(s)
		break
	}
}
