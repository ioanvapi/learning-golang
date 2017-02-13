package main

import (
	"io/ioutil"
	"log"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

func main() {
	bytes, _ := ioutil.ReadFile("../sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes)

	accessToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.qXk8xnU3F8aYKahI8EzStHBYuVDm8BoFDR6JQ7F2oSMgf3B-JrwAZXp8yGO6yqTB7Hd7wAXHZIFsMuBKhlB1RxZv7u5KbAjTXYtR-hYzaEaEhl3DOFVPeamSyEmV-4G1AGDZYnOvAOBJaS4vac3l9aRN4m05WMA_TEbpSTgsUNBoGkUnGdQjgijrc3ecr02aGqtv2MeycC1sPJ7wM69LDtGletrcpOc0abY0pXgkbqT3RXavcOWlObThZGsRb48tQHzCCOkAp1sv8tNOozrtNQOZF1SNbnuvM-ldthAp2usAImziBYh7RB-PnIcoxsR7k7fBcFtfxOVulsG4NlOv_Q"

	jwt, err := jws.ParseJWT([]byte(accessToken))
	if err != nil {
		log.Fatal(err)
	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		log.Fatal(err)
	}
}
