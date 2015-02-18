package main

import "fmt"

func main() {
	// var x string = "Hello World"
	//var x string
	//x = "Hello World"
	//fmt.Println(x)
	var x string
	x = "first "
	fmt.Println(x)
	//x = "second"
	x = x + "second"
	x += " second"
	fmt.Println(x)
	var a string = "hello"
	var b string = "world"
	fmt.Println(a == b)
	var c string = "hello"
	fmt.Println(a == c)
	d := "Hello World"
	fmt.Println(d)
	e := 5
	fmt.Println(e)

	dogName := "Max"
	fmt.Println("My dog's name is", dogName)

	const f string = "Hello World"
	fmt.Println(f)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
