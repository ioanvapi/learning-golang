package main

// 所有不成对向信道存取数据的情况都是死锁? 反例

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
}
