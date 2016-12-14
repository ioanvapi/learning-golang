package main

import (
	"log"
	"sync"
	"time"

	"github.com/golang/groupcache/singleflight"
)

func newDelayReturn(dur time.Duration, n int) func() (interface{}, error) {
	return func() (interface{}, error) {
		time.Sleep(dur)
		return n, nil
	}
}

func main() {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ret, err := g.Do("key", newDelayReturn(time.Second*1, 1))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("key-1 get %v", ret)
		wg.Done()
	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // make sure this is call is later
		ret, err := g.Do("key", newDelayReturn(time.Second*2, 2))
		if err != nil {
			panic(err)
		}
		log.Printf("key-2 get %v", ret)
		wg.Done()
	}()
	wg.Wait()
}
