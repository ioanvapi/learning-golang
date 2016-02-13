package main

func main() {
	// 内置函数 len 和 cap 都返回数组长度 (元素数量)。
	a := [2]int{}
	println(len(a), cap(a)) // 2, 2
}
