package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	MaxEpollEvents = 32
	KB             = 1024
)

func echo(in, out int) {

	var buf [KB]byte
	for {
		nbytes, e := syscall.Read(in, buf[:])
		if nbytes > 0 {
			syscall.Write(out, buf[:nbytes])
		}
		if e != nil {
			break
		}
	}
}

func main() {

	var event syscall.EpollEvent
	var events [MaxEpollEvents]syscall.EpollEvent

	file, e := os.Open("/dev/kmsg")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer file.Close()

	fd := int(file.Fd())
	if e = syscall.SetNonblock(fd, true); e != nil {
		fmt.Println("setnonblock1: ", e)
		os.Exit(1)
	}

	epfd, e := syscall.EpollCreate1(0)
	if e != nil {
		fmt.Println("epoll_create1: ", e)
		os.Exit(1)
	}
	defer syscall.Close(epfd)

	event.Events = syscall.EPOLLIN
	event.Fd = int32(fd)
	if e = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event); e != nil {
		fmt.Println("epoll_ctl: ", e)
		os.Exit(1)
	}

	for {

		nevents, e := syscall.EpollWait(epfd, events[:], -1)
		if e != nil {
			fmt.Println("epoll_wait: ", e)
			break
		}

		for ev := 0; ev < nevents; ev++ {
			go echo(int(events[ev].Fd), syscall.Stdout)
		}

	}

}
