package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"a": 1,
	}
	fmt.Println(m)

	if v, ok := m["a"]; ok {
		// 判断 key 是否存在
		fmt.Println(ok) // true
		println(v)
	}

	println(m["c"])       // 对于不存在的 key，直接返回 \0，不会出错。
	m["b"] = 2            // 新增或修改
	delete(m, "c")        // 删除。如果 key 不存在，不会出错。
	println(len(m))       // 获取键值对数量。cap 无效
	for k, v := range m { // 迭代，可仅返回 key。随机顺序返回，每次都不相同。
		println(k, v)
	}

	// 从 map 中取回的是一个 value 临时复制品，对其成员的修改是没有任何意义的。
	type user struct{ name string }

	m1 := map[int]user{ // 当 map 因扩张而重新哈希时，各键值项存储位置都会发生改变。 因此，map
		1: {"user1"}, // 被设计成 not addressable。 类似 m[1].name 这种期望透过原 value
	} // 指针修改成员的行为自然会被禁止。
	fmt.Println(m1)

	// m1[1].name = "Tom" // Error: cannot assign to m[1].name

	// 正确做法是完整替换 value 或使用指针
	u := m1[1]
	u.name = "Tom"
	m1[1] = u // 替换 value。

	fmt.Println(m1)

	m2 := map[int]*user{
		1: &user{"user1"},
	}

	m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。
	fmt.Println(*m2[1])
}
