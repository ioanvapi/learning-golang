package main

import (
	"fmt"
)

type User struct {
	name string
}

type Manager struct {
	User
	title string
}

type Resource1 struct {
	id int
}

type User1 struct {
	Resource1
	name string
}

type Manager1 struct {
	User1
	title string
}

type Resource2 struct {
	id   int
	name string
}

type Classify2 struct {
	id int
}

type User2 struct {
	Resource2 // Resource.id 与 Classify.id 处于同一层次
	Classify2
	name string // 遮蔽 Resource.name
}

type Resource3 struct {
	id int
}

type User3 struct {
	*Resource3
	// Resource3
	name string
}

func main() {
	// 匿名字段不过是一种语法糖，从根本上说，就是一个与成员类型同名 (不含包名) 的字段。被匿名嵌入的可以是任何类型，当然也包括指针。
	m := Manager{
		User:  User{"Tom"}, // 匿名字段的显式字段名，和类型名相同
		title: "Administrator",
	}
	fmt.Println(m)

	// 可以像普通字段那样访问匿名字段成员，编译器从外向内逐级查找所有层次的匿名字段，直到发现目标或出错。
	var m1 Manager1
	m1.id = 1
	m1.name = "Jack"
	m1.title = "Admin"

	fmt.Println(m1)

	// 外层同名字段会遮蔽嵌入字段成员，相同层次的同名字段也会让编译器无所适从。解决方法是使用显式字段名。
	u := User2{
		Resource2{1, "people"},
		Classify2{100},
		"Jack",
	}
	println(u.name)
	println(u.Resource2.name)
	// println(u.id)
	println(u.Classify2.id)

	u1 := User3{
		&Resource3{1},
		"Admin",
	}

	println(u1.id)
	println(u1.Resource3.id)
}
