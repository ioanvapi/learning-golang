package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

//func main() {
//	finalHandler := http.HandlerFunc(final)
//
//	logFile, err := os.OpenFile("server.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
//	if err != nil {
//		panic(err)
//	}
//
//	http.Handle("/", handlers.LoggingHandler(logFile, finalHandler))
//
//	log.Println("Listening at :3000")
//	http.ListenAndServe(":3000", nil)
//}

func myLoggingHandler(h http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}

func main() {
	finalHandler := http.HandlerFunc(final)

	http.Handle("/", myLoggingHandler(finalHandler))

	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", nil)
}
