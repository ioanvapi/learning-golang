package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {

	handlers := http.NewServeMux()
	handlers.HandleFunc("/1", handler)
	server := &http.Server{Addr: ":8888", Handler: handlers}

	// test 1
	// go server.ListenAndServe()

	// handlers.HandleFunc("/2", handler)

	// select {}

	// test 2
	handlers.HandleFunc("/2", handler)
	server.ListenAndServe()

}
