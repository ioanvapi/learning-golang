package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(3 * time.Second)
		close(jobs)
	}()
	go func() {
		for i := 0; ; i++ {
			jobs <- i
			fmt.Println("produce:", i)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
		}
	}()
	wg.Wait()
}
