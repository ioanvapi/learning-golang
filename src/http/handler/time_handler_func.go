package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	// Convert the timeHandler function to a HandlerFunc type
	//th := http.HandlerFunc(timeHandler)

	// And add it to the ServeMux
	//mux.Handle("/time", th)

	mux.HandleFunc("/time", timeHandler)

	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", mux)
}
