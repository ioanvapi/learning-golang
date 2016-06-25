package main

import (
	"encoding/json"
	"os"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}

func main() {
	book := Book{
		Id:   1,
		Name: "golang book",
		Categories: []Category{
			{Id: 3, Name: "golang"},
			{Id: 4, Name: "tutorial"},
		},
	}

	json.NewEncoder(os.Stdout).Encode(&book)
}
