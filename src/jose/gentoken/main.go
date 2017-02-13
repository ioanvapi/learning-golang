package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

func main() {
	bytes, _ := ioutil.ReadFile("../sample_key.priv")

	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(time.Duration(10) * time.Second))

	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(rsaPrivate)
	fmt.Printf("jwt serial: %s\n", b)
}
