package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string ",omitempty"
}

func main() {
	data, err := bson.Marshal(&Person{Name:"Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", data)
}
