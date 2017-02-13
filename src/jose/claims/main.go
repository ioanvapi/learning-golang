package main

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/jws"
)

func main() {
	// expires in 10 seconds
	now := time.Now()
	expires := now.Add(time.Duration(10) * time.Second)

	claims := jws.Claims{}
	claims.SetExpiration(expires)
	claims.SetIssuedAt(now)

	fmt.Printf("claims: %v\n", claims)
}
