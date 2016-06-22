package main

import (
	"net/http"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"
)

func main() {
	a := app.New()

	a.Use(logger.New())
	a.Use(serve.New("examples"))

	a.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))

	a.Listen(":3000")
}
