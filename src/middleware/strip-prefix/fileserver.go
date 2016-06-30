package main

import "net/http"

func main() {
	// http://localhost:8888/static
	// http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	// http://localhost:8888
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8888", nil)
}
