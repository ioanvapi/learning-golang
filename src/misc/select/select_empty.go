package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("111111")
		time.Sleep(5 * time.Second)
		fmt.Println("111111 exit")
	}()

	go func() {
		fmt.Println("222222")
		time.Sleep(10 * time.Second)
		fmt.Println("222222 exit")
	}()

	select {}
}
