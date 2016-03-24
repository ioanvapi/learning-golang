package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	jsonParsed, _ := gabs.ParseJSON([]byte(`{"array":[ "first", "second", "third" ]}`))

	// S is shorthand for Search
	children, _ := jsonParsed.S("array").Children()
	for _, child := range children {
		fmt.Println(child.Data().(string))
	}
}
