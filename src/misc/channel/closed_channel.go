package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		// 第二个表示 channel 的启用状态，当前是 false，表示 channel 被关闭
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
