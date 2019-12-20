package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	// TODO: Actually run an image!
	if os.Args[1] != "run" || os.Args[2] != "codecraftersio/docker-challenge-1" {
		fmt.Printf("Expected 'run <image> <command>' as the command! args: %q\n", os.Args[1:])
		os.Exit(1)
	}

	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Printf("TempDir Error: %v", err)
		os.Exit(1)
	}

	if err = os.MkdirAll(tempDir+"/usr/bin", os.ModeDir); err != nil {
		fmt.Printf("Mkdir error: %v", err)
		os.Exit(1)
	}

	src := "/usr/bin/docker-explorer"
	dest := tempDir + "/usr/bin/docker-explorer"

	if err := copyExecutable(src, dest); err != nil {
		fmt.Printf("Copy Executable Error: %v", err)
		os.Exit(1)
	}

	forkAttr := syscall.ProcAttr{
		Env: os.Environ(),
		Sys: &syscall.SysProcAttr{
			Chroot: tempDir,
		},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
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

func copyExecutable(src string, dest string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}
