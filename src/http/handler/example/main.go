package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://httpbin.org/get", http.StatusTemporaryRedirect)
	mux.Handle("/foo", rh)

	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", mux)
}
