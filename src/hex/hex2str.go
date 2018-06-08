package main

import (
	"log"
	"encoding/hex"
	"crypto/md5"
	"fmt"
)

func main() {

	{
		// hex string to []byte
		hexStr := "fee9ecaadafeee72d2eb66a0bd344cdd"
		data, err := hex.DecodeString(hexStr)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("hex decode data: %v", data)
	}

	{
		// []byte to hex string
		data := "test string"
		// md5.Sum() return a byte array
		h := md5.Sum([]byte(data))

		// with "%x" format byte array into hex string
		hexStr := fmt.Sprintf("%x", h)
		log.Printf("hex str: %v", hexStr)
	}

}