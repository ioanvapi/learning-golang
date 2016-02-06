package main

import (
	"fmt"
	"flag"
)

// bind the flag to a variable
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagvar", 1234, "help message for flagvar")
}

func main() {
	var flagint = flag.Int("flagint", 4321, "help message for flagint")
	flag.Parse()
	fmt.Println("flagint: ", *flagint)
	fmt.Println("flagvar: ", flagvar)
}