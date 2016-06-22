package main

import (
	"fmt"
	"net/http"
)

type AppendMiddleware struct {
	handler http.Handler
}

func (a *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	w.Write([]byte("Middleware says hello."))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	appendMid := new(AppendMiddleware)
	appendMid.handler = http.HandlerFunc(helloHandler)
	http.ListenAndServe(":8081", appendMid)
}
