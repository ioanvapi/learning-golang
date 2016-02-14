package main

import (
	"fmt"
)

// 接口对象由接口表 (interface table) 指针和数据指针组成

type User struct {
	id   int
	name string
}

// 数据指针持有的是目标对象的只读复制品，复制完整对象或指针

func main() {
	u := User{1, "Tom"}
	var i interface{} = u

	u.id = 2
	u.name = "Jack"

	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))
	// 接口转型返回临时对象，只有使用指针才能修改其状态
	var vi, pi interface{} = u, &u

	// vi.(User).name = "Jack"
	pi.(*User).name = "Jack"

	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))
}
