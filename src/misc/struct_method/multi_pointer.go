package main

type X struct{}

func (*X) test() {
	println("X.test")
}

func main() {
	p := &X{}
	p.test()

	// 从 1.4 开始，不再支持多级指针查找方法成员
	// (&p).test()
}
