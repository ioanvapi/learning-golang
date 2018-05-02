package main

import (
	"fmt"

	"github.com/eoscanada/eos-go/ecc"
)

func main() {
	wif, err := ecc.NewRandomPrivateKey()
	if err != nil {
		fmt.Println(err)
	}
	privKey, err := ecc.NewPrivateKey(wif.String())
	if err != nil {
		fmt.Println(err)
	}

	pubKey := privKey.PublicKey()

	fmt.Printf("Private: %v\nPublic: %s\n", wif.String(), pubKey.String())
}
