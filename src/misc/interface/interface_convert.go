package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
	// 利用类型推断，可判断接口对象是否某个具体的接口或类型
	var o interface{} = &User{1, "Tom"}
	fmt.Println(o)

	if i, ok := o.(fmt.Stringer); ok { // ok-idiom
		fmt.Println(i)
	}

	u := o.(*User)
	// u := o.(User) // panic: interface is *main.User, not main.User
	fmt.Println(u)

	// 还可用 switch 做批量类型判断，不支持 fallthrough
	var o1 interface{} = &User{1, "Tom"}

	switch v := o1.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println("interface", v)
	case func() string: // func
		fmt.Println("func", v())
	case *User: // *struct
		fmt.Printf("*struct, %d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}
