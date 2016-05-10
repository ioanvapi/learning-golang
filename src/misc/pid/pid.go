package main

import (
	"fmt"
	"os"
)

func main() {
	ppid := os.Getppid()
	pid := os.Getpid()
	done := make(chan bool, 1)

	go func() {
		sub_ppid := os.Getppid()
		sub_pid := os.Getpid()

		fmt.Println("sub ppid: ", sub_ppid)
		fmt.Println("sub pid: ", sub_pid)
		done <- true
	}()

	fmt.Println("main ppid: ", ppid)
	fmt.Println("main pid: ", pid)
	<-done
	fmt.Println("exiting")
}
