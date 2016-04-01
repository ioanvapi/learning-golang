package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(url string) []byte {
	res, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func main() {
	clientID := "<client id>"
	body := fetch("https://api.instagram.com/v1/media/popular?client_id=" + clientID)
	fmt.Printf("%s", body)
}
