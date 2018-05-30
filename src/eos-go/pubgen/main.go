package main

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-go/ecc"
)

func main() {
	privKey, err := ecc.NewPrivateKey(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	pubKey := privKey.PublicKey()

	fmt.Printf("Private: %v\nPublic: %s\n", os.Args[1], pubKey.String())
}
