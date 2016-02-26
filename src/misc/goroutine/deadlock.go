package main

import "fmt"

// 只在单一的goroutine里操作无缓冲信道，一定死锁
func main() {
	ch := make(chan int)
	// <-ch // 阻塞main goroutine, 信道ch被锁
	ch <- 0                                // both cause deadlock, 0流入信道，堵塞当前线, 没人取走数据信道不会打开
	fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
}
