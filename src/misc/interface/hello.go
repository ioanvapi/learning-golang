package main

import (
	"fmt"
)

type s struct { // 定义一个s类型, 有一个属性i是int的
	i int
}

func (this *s) Get() int { // Get方法获得i属性
	return this.i
}

func (this *s) Put(v int) { // Put方法设置i属性
	this.i = v
}

type I interface { // 定义一个接口类型, 里面有Get方法与Put方法
	Get() int
	Put(int)
}

func f(my I) { // 这里的my保存了接口类型的值，因为s实现了I，所以传递的my虽然是个I类型，但是可以当作s类型来使用
	my.Put(999)
	fmt.Println(my.Get())
}

func main() {
	var S s // 申请一个S变量, 他是s类型的值
	f(&S)
}
