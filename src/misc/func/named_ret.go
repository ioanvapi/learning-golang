package main

import (
	"fmt"
)

// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
func add(x, y int) (z int) {
	z = x + y
	return
}

// 命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
func add1(x, y int) (z int) {
	{
		// 不能在一个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return // error
		return z // 必须显式返回
	}
}

// 命名返回参数允许 defer 延迟调用通过闭包读取和修改。
func add2(x, y int) (z int) {
	defer func() {
		z += 100
	}()

	z = x + y
	return
}

// 显式 return 返回前，会先修改命名返回参数
func add3(x, y int) (z int) {
	defer func() {
		fmt.Println(z) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执⾏行顺序: (z = z + 200) -> (call defer) -> (ret)
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add1(1, 2))
	fmt.Println(add2(1, 2))
	fmt.Println(add3(1, 2))
}
