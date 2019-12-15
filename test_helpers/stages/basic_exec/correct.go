package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// TODO: Actually run an image!
	if os.Args[1] != "run" || os.Args[2] != "codecraftersio/docker-challenge-1" {
		fmt.Printf("Expected 'run <image> <command>' as the command! args: %q\n", os.Args[1:])
		os.Exit(1)
	}

	if err := syscall.Exec("/usr/bin/"+os.Args[3], os.Args[3:], os.Environ()); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
