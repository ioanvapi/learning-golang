package main

import (
	"log"
	"net/http"

	ohttp "github.com/ossrs/go-oryx-lib/http"
)

func main() {
	ohttp.Server = "akserver"
	data := "data string"
	fn := ohttp.Data(nil, data).(http.HandlerFunc)
	http.HandleFunc("/data_string", fn)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
