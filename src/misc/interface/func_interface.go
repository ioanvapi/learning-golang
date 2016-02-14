package main

type Tester interface {
	Do()
}

type FuncDo func()

func (self FuncDo) Do() {
	self()
}

func main() {
	// 让函数直接 "实现" 接口能省不少事
	var t Tester = FuncDo(func() { println("Hello, World!") })
	t.Do()
}
