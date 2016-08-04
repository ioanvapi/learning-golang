package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("do something")
				time.Sleep(1 * time.Second)
				// other works ...
			}
		}
	}()

	time.Sleep(2 * time.Second)
	// signal all relevant goroutines
	close(done)
}
