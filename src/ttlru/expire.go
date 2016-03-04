package main

import (
  "fmt"
  "time"
  "github.com/zvelo/ttlru"
)

func main() {
  cache := ttlru.New(10, 3 * time.Second)
  fmt.Println(cache.Cap(), cache.Len())

  cache.Set("foo", "bar")

  val, err := cache.Get("foo")
  fmt.Println(val, err)

  time.Sleep(2 * time.Second)

  val, err = cache.Get("foo")
  fmt.Println(val, err)

  time.Sleep(4 * time.Second)

  val, err = cache.Get("foo")
  fmt.Println(val, err)
}
