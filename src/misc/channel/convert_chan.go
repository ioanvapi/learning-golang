package main

import (
	"fmt"
	"time"
)

func stopTimeout(t <-chan time.Time) chan bool {
	b := make(chan bool, 1)
	select {
	case <-t:
		b <- true
		return b
	}
}

func main() {
	c := stopTimeout(time.After(2 * time.Second))
	d := <-c
	fmt.Println(d)
}
