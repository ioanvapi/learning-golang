package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type JsonHolder struct {
	Data map[string]string `json:"data"`
}

func main() {
	var m, s string
	flag.StringVar(&m, "m", "", "map data")
	flag.StringVar(&s, "s", "", "slice data")
	flag.Parse()

	// map
	jsonData := fmt.Sprintf("{\"data\":%s}", m)
	jsonHolder := JsonHolder{}
	json.Unmarshal([]byte(jsonData), &jsonHolder)

	mp := jsonHolder.Data

	fmt.Printf("%#v\n", mp)

	// map 2
	var mp2 map[string]string
	json.Unmarshal([]byte(m), &mp2)
	fmt.Printf("%#v\n", mp2)

	// slice
	var sl []string
	json.Unmarshal([]byte(s), &sl)
	fmt.Printf("%#v\n", sl)
}
