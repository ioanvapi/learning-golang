package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	m := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"scan": func(e *fsm.Event) {
				fmt.Println("after_scan: " + e.FSM.Current())
			},
			"working": func(e *fsm.Event) {
				fmt.Println("working: " + e.FSM.Current())
			},
			"situation": func(e *fsm.Event) {
				fmt.Println("situation: " + e.FSM.Current())
			},
			"finish": func(e *fsm.Event) {
				fmt.Println("finish: " + e.FSM.Current())
			},
		},
	)

	fmt.Println(m.Current())

	err := m.Event("scan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + m.Current())

	err = m.Event("working")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + m.Current())

	err = m.Event("situation")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + m.Current())

	err = m.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + m.Current())

}
