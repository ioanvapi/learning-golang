package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	finalHandler := http.HandlerFunc(final)
	authHandler := httpauth.SimpleBasicAuth("username", "password")

	http.Handle("/", authHandler(finalHandler))

	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", nil)
}
