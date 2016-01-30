package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    `how much there are`
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	var reflectType reflect.Type = reflect.TypeOf(tt)
	var ixField reflect.StructField
	for i := 0; i < 3; i++ {
		ixField = reflectType.Field(i)
		fmt.Printf("%s\n", ixField.Tag)
	}
}
