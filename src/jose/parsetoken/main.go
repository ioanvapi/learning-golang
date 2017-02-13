package main

import (
	"fmt"
	"net/http"

	"github.com/SermoDigital/jose/jws"
)

func ParseTokenHandler(rw http.ResponseWriter, r *http.Request) {
	j, err := jws.ParseFromHeader(r, jws.General)

	fmt.Printf("JWS: %v, err: %v", j, err)

	// Validate token here...
	// j.Validate(rsaPublic, crypto.SigningMethodRS256)
}

func main() {
	http.HandleFunc("/", ParseTokenHandler)
	http.ListenAndServe(":3000", nil)
}
