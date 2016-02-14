package main

import "fmt"

// 面向对象三大特征里，Go 仅支持封装，尽管匿名字段的内存布局和行为类似继承。没有 class 关键字，没有继承、多态等等。
type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User{1, "Tom"}, "Admin"}
	// var u User = m      // 没有继承，自然也不会有多态
	var u User = m.User // 同类型拷⻉
	fmt.Println(m, u)
}
