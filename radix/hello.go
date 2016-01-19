package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 10*time.Second)
	errHndlr(err)
	defer c.Close()

	// r := c.Cmd("select", 8)
	// errHndlr(r.Err)

	r := c.Cmd("flushdb")
	errHndlr(r.Err)

	r = c.Cmd("echo", "Hello world!")
	errHndlr(r.Err)

	r = c.Cmd("set", "mykey0", "myval0")
	errHndlr(r.Err)
}
