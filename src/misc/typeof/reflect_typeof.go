package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	ppid := os.Getppid()
	fmt.Println("type ppid: ", reflect.TypeOf(ppid))
}
