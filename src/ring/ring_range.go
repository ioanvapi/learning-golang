package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	coffee := []string{"kenya", "guatemala", "ethiopia"}

	// create a ring and populate it with some values
	r := ring.New(len(coffee))
	fmt.Println("r.Len()", r.Len())
	for i := 0; i < r.Len(); i++ {
		r.Value = coffee[i]
		r = r.Next()
		// r = r.Move(1)
	}

	// print all values of the ring, easy done with ring.Do()
	r.Do(func(x interface{}) {
		fmt.Println(x)
	})

	for i := 0; i < r.Len()+4; i++ {
		fmt.Println("r.Value", i, r.Value)
		// r = r.Next()
		r = r.Move(1)
	}

	// .. or each one by one.
	for _ = range time.Tick(time.Second * 1) {
		r = r.Next()
		fmt.Println(r.Value)
	}
}
