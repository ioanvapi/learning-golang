package main

import (
	"encoding/json"
	"fmt"
)

const text = `
{
	"foo": "hello",
	"bar": ["golang!", "sdfsdfsdf"],
  "ffo": {
    "aaa": "aaaa",
    "bbb": "bbb"
  }
}
`

func main() {
	var m = make(map[string]interface{})
	json.Unmarshal([]byte(text), &m)
	fmt.Println(m)
}
