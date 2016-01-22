package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

func main() {
	//初始化
	person := Person{"Tom", 30, "tom@gmail.com"}
	person = Person{name:"Tom", age: 30, email:"tom@gmail.com"}

	fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

	pPerson := &person

	fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}

	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}
}