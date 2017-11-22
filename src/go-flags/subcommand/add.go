package main

import (
	"fmt"
)

type AddCmd struct {
	All bool `short:"a" long:"all" description:"Add all files"`
}

func (ac *AddCmd) Execute(args []string) error {
	fmt.Printf("Adding (all=%v): %#v\n", ac.All, args)
	return nil
}

var addCmd AddCmd

func init() {
	parser.AddCommand("add", "Add a file", "The add command adds a file to the repository. Use -a to add all files.", &addCmd)
}
