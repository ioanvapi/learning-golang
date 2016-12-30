package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	m := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open_event", Src: []string{"closed"}, Dst: "open"},
			{Name: "close_event", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(m.Current())

	err := m.Event("open_event")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.Current())

	err = m.Event("close_event")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.Current())
}
