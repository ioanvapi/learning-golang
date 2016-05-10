package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds
	c := cache.New(3*time.Second, 2*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	time.Sleep(1 * time.Second)

	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	time.Sleep(1 * time.Second)

	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	time.Sleep(1 * time.Second)

	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	time.Sleep(1 * time.Second)

	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
