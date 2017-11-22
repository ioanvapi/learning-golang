package main

import (
	"fmt"
)

type RmCmd struct {
	Force bool `short:"f" long:"force" description:"Force removal of files"`
}

func (rc *RmCmd) Execute(args []string) error {
	fmt.Printf("Removing (force=%v): %#v\n", rc.Force, args)
	return nil
}

var rmCmd RmCmd

func init() {
	parser.AddCommand("rm", "Remove a file", "The rm command removes a file to the repository. Use -f to force removal of files.", &rmCmd)
}
