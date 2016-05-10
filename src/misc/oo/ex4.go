package main

import (
	"fmt"
)

type Rect struct {
	width  int
	height int
}

type Rects []Rect

func (r *Rect) Area() int {
	return r.width * r.height
}

// interface declaration
type Shaper interface {
	Area() int
}

// using interface as param type
func Describe(s Shaper) {
	fmt.Println("Area is: ", s.Area())
}

func main() {
	r := &Rect{width: 10, height: 5}
	x := &Rect{width: 7, height: 10}
	//	rs := &Rects{r, x}
	Describe(r)
	Describe(x)
	//	Describe(rs)
}
