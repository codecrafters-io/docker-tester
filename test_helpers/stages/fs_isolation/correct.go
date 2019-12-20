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

	if err := syscall.Chroot(tempDir); err != nil {
		fmt.Printf("Chroot Error: %v", err)
		os.Exit(1)
	}

	if err := syscall.Exec(os.Args[3], os.Args[3:], os.Environ()); err != nil {
		fmt.Printf("Exec Error: %v", err)
		os.Exit(1)
	}
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
