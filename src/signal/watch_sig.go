package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CreateInterrupt Creates a signal interrupt so that upon a Ctrl-C (as well as some others)
// Print(div) will be called and then the process will be exited
func CreateInterrupt() {
	go func() {
		log.Println("Waiting for signal")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
		<-c
		go func() {
			<-c
			os.Exit(1)
		}()
		log.Println("Got SIG")
		os.Exit(0)
	}()
}

func main() {
	CreateInterrupt()
	time.Sleep(10 * time.Second)
}
