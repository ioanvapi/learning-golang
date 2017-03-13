package main

import "fmt"

func main() {
	b, err := msgpack.Marshal(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)

	var out bool
	err = msgpack.Unmarshal([]byte{0xc3}, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: []byte{0xc3}
	// true
}
