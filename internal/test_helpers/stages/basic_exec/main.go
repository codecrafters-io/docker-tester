package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// TODO: Actually run an image!
	if os.Args[1] != "run" || os.Args[2] != "codecraftersio/docker-challenge" {
		fmt.Printf("Expected 'run <image> <command>' as the command! args: %q\n", os.Args[1:])
		os.Exit(1)
	}

	forkAttr := syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdout.Fd(), os.Stdout.Fd(), os.Stdout.Fd()},
	}

	pid, err := syscall.ForkExec(os.Args[3], os.Args[3:], &forkAttr)
	if err != nil {
		fmt.Printf("Fork Error: %v", err)
		os.Exit(1)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("FindProcess Error: %v", err)
		os.Exit(1)
	}

	state, err := process.Wait()
	if err != nil {
		fmt.Printf("ProcessWait Error: %v", err)
		os.Exit(1)
	}

	os.Exit(state.ExitCode())
}
