package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// TODO: Actually run an image!
	if os.Args[1] != "run" || os.Args[2] != "alpine" {
		fmt.Println("Expected 'run alpine <command>' as the command!")
	}

	if err := syscall.Exec("/usr/bin/"+os.Args[3], os.Args[3:], os.Environ()); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
