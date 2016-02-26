package main

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 1
		quit <- 0
	}()

	<-c // 取走c的数据！
	<-quit
}
