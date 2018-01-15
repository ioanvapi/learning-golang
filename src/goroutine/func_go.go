package main

import (
	"fmt"
	"runtime"
)

type ClassA struct {
	a int
}

func (c *ClassA) f() int {
	c.a = 0
	go func() {
		c.a = 1
		//time.Sleep(1*time.Second)
	}()
	runtime.Gosched()
	return c.a
}

func main() {
	c := ClassA{}
	fmt.Println(c.f())
	//time.Sleep(2*time.Second)
	//fmt.Println(c.a)
}
