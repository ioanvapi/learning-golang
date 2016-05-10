package main

import "fmt"

func main() {
	v, e := multi_ret("one")
	fmt.Println(v, e) //输出 1 true

	v, e = multi_ret("four")
	fmt.Println(v, e) //输出 0 false

	//通常的用法(注意分号后有e)
	if v, e = multi_ret("four"); e {
		// 正常返回
		fmt.Println(v, e)
	} else {
		// 出错返回
		fmt.Println(v, e)
	}
}

func multi_ret(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int

	val, err = m[key]

	return val, err
}
