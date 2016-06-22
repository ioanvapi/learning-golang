package main

import (
	"fmt"
	"net/http"

	"github.com/gohttp/app"
)

func main() {
	app := app.New()

	app.Get("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "foo")
	}))

	app.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "bar")
	})

	app.Listen(":3000")
}
