package main

import (
	"fmt"
)

// 值类型，赋值和传参会复制全部内容。可用 "_" 定义补位字段，支持指向自身类型的指针成员。

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

type User struct {
	name string
	age  int
}

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	fmt.Println(n1, n2)

	// 顺序初始化必须包含全部字段，否则会出错
	u1 := User{"Tom", 20}
	// u2 := User{"Tom"}
	fmt.Println(u1)

	f := File{
		name: "test.txt",
		size: 1025,
		// attr: {0755, 1},
	}

	f.attr.owner = 1
	f.attr.perm = 0755

	fmt.Println(f)

	var attr = struct {
		perm  int
		owner int
	}{2, 0755}

	f.attr = attr
	fmt.Println(f)

	// 支持 "=="、"!=" 相等操作符，可用作 map 键类型。
	m := map[User]int{
		User{"bob", 19}: 100,
	}
	fmt.Println(m)

	// 可定义字段标签，用反射读取。标签是类型的组成部分。
	var u3 struct {
		name string "username"
	}
	var u4 struct {
		name string
	}
	// u2 = u1
	fmt.Println(u3, u4)

	// 空结构 "节省" 内存，比如用来实现 set 数据结构，或者实现没有 "状态" 只有方法的 "静态类"。
	var null struct{}

	set := make(map[string]struct{})
	set["a"] = null
	fmt.Println(set)
}
