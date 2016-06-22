package main

import (
	"fmt"
	"net/http"
)

func SingleHost(handler http.Handler, allowedHost string) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
			fmt.Fprintf(w, "you got 403")
		}
	}
	return http.HandlerFunc(ourFunc)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	singleHosted := SingleHost(http.HandlerFunc(helloHandler), "localhost:8081")
	http.ListenAndServe(":8081", singleHosted)
}
