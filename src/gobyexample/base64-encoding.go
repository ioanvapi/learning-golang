package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Standard encoded: %v\n", sEnc)

	sEnc1 := b64.RawStdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Standard with padding encoded: %v\n", sEnc1)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Printf("Standard dencoded: %v\n\n", string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Printf("URL-compatible encoded: %v\n", uEnc)

	uEnc1 := b64.RawURLEncoding.EncodeToString([]byte(data))
	fmt.Printf("URL-compatible with padding encoded: %v\n", uEnc1)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Printf("URL-compatible dencoded: %v\n", string(uDec))
}
