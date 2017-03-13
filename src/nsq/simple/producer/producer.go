package main

import "log"

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("10.0.3.126:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
