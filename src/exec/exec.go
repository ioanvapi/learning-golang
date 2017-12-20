package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// os.StartProcess
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!\n", err)
		os.Exit(1)
	}

	fmt.Printf("The ls process info is %+v\n", pid)

	// 2nd example: show all processes
	pid, err = os.StartProcess("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!\n", err)
		os.Exit(1)
	}

	fmt.Printf("The ps process info is %+v\n", pid)

	// exec.Run
	cmd := exec.Command("ls", "-l")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!\n", err)
		os.Exit(1)
	}

}
