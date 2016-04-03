package main

import (
	"os"
	"strings"
)

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	_ = s
	// fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	_ = s
	// fmt.Println(s)
}

func echo3(args []string) {
	s := strings.Join(args[1:], " ")
	_ = s
	// fmt.Println(s)
}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}
