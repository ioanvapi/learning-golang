package main

import (
	"sort"
	"fmt"
)

func sorted(lsa []float64) []float64 {
	lsb := make([]float64, len(lsa))
	copy(lsb, lsa)
	sort.Float64s(lsb)
	return lsb
}

func main() {
	lsa := []float64{1.0, 3.0, 2.0}
	lsb := sorted(lsa)
	fmt.Println("lsa: ", lsa)
	fmt.Println("lsb: ", lsb)
}
