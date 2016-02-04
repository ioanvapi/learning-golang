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

	r := c.Cmd("select", 0)
	errHndlr(r.Err)

	r = c.Cmd("PING")
	errHndlr(r.Err)
	s, _ := r.Str()
	fmt.Println("PING get: ", s)

	r = c.Cmd("flushdb")
	errHndlr(r.Err)

	r = c.Cmd("echo", "Hello world!")
	errHndlr(r.Err)

	r1, err1 := c.Cmd("echo", "Hello world!").Str()
	errHndlr(err1)
	fmt.Println("echo: ", r1)

	r = c.Cmd("echo", "Hello world!")
	errHndlr(r.Err)
	s, _ = r.Str()
	fmt.Println("echo: ", s)

	r = c.Cmd("set", "mykey0", "myval0")
	errHndlr(r.Err)

	r = c.Cmd("MSET", "key1", "val1", "key2", "val2", "key3", "val3")
	errHndlr(r.Err)

	// Array Replies
	r = c.Cmd("MGET", "key1", "key2", "key3")
	errHndlr(r.Err)

	l, _ := r.List()
	for _, elemStr := range l {
		fmt.Println(elemStr)
	}

	elems, err := r.Array()
	for i := range elems {
		elemStr, _ := elems[i].Str()
		fmt.Println(elemStr)
	}

	r = c.Cmd("SET", "foo", "bar")
	errHndlr(r.Err)

	r = c.Cmd("SET", "baz", "bazval")
	errHndlr(r.Err)

	// Pipelining
	c.PipeAppend("GET", "foo")
	c.PipeAppend("SET", "bar", "foo")
	c.PipeAppend("DEL", "baz")

	// Read GET foo reply
	foo, err := c.PipeResp().Str()
	errHndlr(err)
	fmt.Println("GET foo: ", foo)

	// Read SET bar foo reply
	if err := c.PipeResp().Err; err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	r = c.Cmd("GET", "bar")
	errHndlr(r.Err)
	s, _ = r.Str()
	fmt.Println("GET bar: ", s)

	// Read DEL baz reply
	if err := c.PipeResp().Err; err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	// Flattening
	c.Cmd("HMSET", "myhash", "key1", "val1", "key2", "val2")
	c.Cmd("HMSET", "myhash", []string{"key1", "val1", "key2", "val2"})
	c.Cmd("HMSET", "myhash", map[string]string{
		"key1": "val1",
		"key2": "val2",
	})
	c.Cmd("HMSET", "myhash", [][]string{
		[]string{"key1", "val1"},
		[]string{"key2", "val2"},
	})
}
