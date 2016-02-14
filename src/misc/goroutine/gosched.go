package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 和协程 yield 作用类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执行

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			fmt.Println(i)

			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello, World!")
	}()

	wg.Wait()
}
