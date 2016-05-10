package main

import (
	"fmt"
)

// type declaration
type Rect struct {
	width  int
	height int
}

// type declaration (slice)
type Rects []Rect

// declaring a method
func (r *Rect) Area() int {
	return r.width * r.height
}

func (rs Rects) Area() int {
	var a int
	for _, r := range rs {
		a += r.Area()
	}

	return a
}

func main() {
	r := Rect{
		width:  10,
		height: 5,
	}

	x := Rect{
		width:  7,
		height: 10,
	}

	rs := Rects{r, x}
	fmt.Println("r's area: ", r.Area())
	fmt.Println("x's area: ", x.Area())
	fmt.Println("total area: ", rs.Area())

}
