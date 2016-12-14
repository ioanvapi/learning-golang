package main

import (
	"log"

	"github.com/golang/groupcache/consistenthash"
)

func main() {
	c := consistenthash.New(70, nil)
	c.Add("A", "B", "C", "D", "E")
	for _, key := range []string{"what", "nice", "what", "nice", "good", "yes!"} {
		log.Printf("%s -> %s\n", key, c.Get(key))
	}
}
