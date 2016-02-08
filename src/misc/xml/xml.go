package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// 人物档案
type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"年龄"`
}

func main() {
	p := person{Name: "bob", Age: 18}

	var data []byte
	var err error

	if data, err = xml.Marshal(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(data))

	if data, err = xml.MarshalIndent(p, "#", " "); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	p2 := new(person)

	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(p2)

}
