package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type ModifierMiddleware struct {
	handler http.Handler
}

func (m *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	// passing a ResponseRecorder instead of the original RW
	m.handler.ServeHTTP(rec, r)
	// after this finishes, we have the response recorded
	// and can modify it before copying it to the original RW

	// we copy the original headers first
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	// and set an additional one
	w.Header().Set("X-We-Modified-This", "Yup")
	// only then the status code, as this call writes out the headers
	w.WriteHeader(418)

	// The body hasn't been written (to the real RW) yet,
	// so we can prepend some data.
	data := []byte("Middleware says hello again. ")

	// But the Content-Length might have been set already,
	// we should modify it by adding the length
	// of our own data.
	// Ignoring the error is fine here:
	// if Content-Length is empty or otherwise invalid,
	// Atoi() will return zero,
	// which is just what we'd want in that case.
	clen, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	clen += len(data)
	r.Header.Set("Content-Length", strconv.Itoa(clen))

	// finally, write out our data
	w.Write(data)
	// then write out the original body
	w.Write(rec.Body.Bytes())
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	modifyMid := new(ModifierMiddleware)
	modifyMid.handler = http.HandlerFunc(helloHandler)
	http.ListenAndServe(":8081", modifyMid)
}
