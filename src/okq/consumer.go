package main

import (
	"log"

	"github.com/mediocregopher/okq-go/okq"
)

func main() {
	cl := okq.New("127.0.0.1:4777")
	defer cl.Close()

	fn := func(e *okq.Event) bool {
		log.Printf("event received on %s: %s\n", e.Queue, e.Contents)
		return true
	}

	for {
		err := cl.Consumer(fn, nil, "super-queue")
		if err != nil {
			log.Printf("Error received from consumer: %s\n", err)
		}
	}
}
