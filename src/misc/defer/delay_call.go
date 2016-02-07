package main

import (
	"os"
)

func test() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	defer f.Close() // 注册调用，而不是注册函数。必须提供参数，哪怕为空。

	f.WriteString("Hello, World!\n")
	return nil
}

func main() {
	test()
}
