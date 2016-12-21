// go run main.go
package main

// #include "lib.h"
// #cgo LDFLAGS: ${SRCDIR}/lib.a -lstdc++
import (
	"C"
	"fmt"
)

func main() {
	fmt.Println("version is", C.get_version())
}
