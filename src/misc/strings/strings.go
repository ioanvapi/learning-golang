package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"

	// 是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))

	// 索引, base 0
	fmt.Println(strings.Index(s, "o"))

	ss := "1#2#345"

	// 切割字符串
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	// 合并字符串
	fmt.Println(strings.Join(splitedStr, "#"))
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))
}
