package main

import (
	"log"

	"github.com/golang/groupcache/lru"
)

func main() {
	cache := lru.New(2)
	cache.Add("x", "x0")
	cache.Add("y", "y0")
	yval, ok := cache.Get("y")
	if ok {
		log.Printf("y is %v", yval)
	}
	cache.Add("z", "z0")

	log.Printf("cache length is %v", cache.Len())
	_, ok = cache.Get("x")
	if !ok {
		log.Println("x key was evicted")
	}
}
