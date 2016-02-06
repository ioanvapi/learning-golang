package main

// 编译器会将未使用的局部变量当做错误

var s string // 全局变量没问题

func main() {
	i := 0 // Error: i declared and not used。(可使用 "_ = i" 规避)
	_ = i
}
