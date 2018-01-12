package main

import (
	"fmt"

	"github.com/jackpal/gateway"
)

func main() {
	fmt.Println(gateway.DiscoverGateway())
}