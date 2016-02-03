package main

import "fmt"

func Greeting(prefix string, who ...string) {
	for _, name := range who {
		fmt.Println(prefix, name)
	}
}

func main() {
	users := []string{"tony", "Alice", "Mike"}

	Greeting("hello", "tony", "Alice", "Mike")

	Greeting("bye", users ...)
	Greeting("bye", users...) // both ok
}
