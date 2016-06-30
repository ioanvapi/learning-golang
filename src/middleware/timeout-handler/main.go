package main

import (
	"net/http"
	"time"
)

const (
	timeout    = time.Duration(1 * time.Second)
	timeoutMsg = "your request has timed out"
)

func myTimeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, timeout, timeoutMsg)
}

func main() {
	indexHandler := http.HandlerFunc(index)

	http.Handle("/", myTimeoutHandler(indexHandler))
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
