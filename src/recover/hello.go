package main

import (
	"log"
	"net/http"

	"github.com/dre1080/recover"
)

var myPanicHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	panic("You should not have a handler that just panics")
})

func main() {
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})

	// recoveryWithDefaults := recovery.New(nil)

	app := recovery(myPanicHandler)
	log.Println("Listening at :3000")
	http.ListenAndServe("0.0.0.0:3000", app)
}
