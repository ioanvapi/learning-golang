package main

import (
	"fmt"
	"net/http"
)

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Println("host: ", host)
	fmt.Println("allowedHost: ", s.allowedHost)
	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "you got 403")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	singleHosted := NewSingleHost(http.HandlerFunc(helloHandler), "localhost:8080")
	http.ListenAndServe(":8080", singleHosted)
}
