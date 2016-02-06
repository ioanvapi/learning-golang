package main

import (
	"fmt"
)

type square struct {
	r int
}

type circle struct {
	r int
}

func (s square) area() int {
	return s.r * s.r
}

func (c circle) area() int {
	return c.r * 3
}

func main() {
	s := square{1}
	c := circle{1}
	fmt.Println(s, c, s.area()+c.area())
}
