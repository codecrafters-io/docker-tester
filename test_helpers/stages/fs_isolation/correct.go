package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	// TODO: Actually run an image!
	if os.Args[1] != "run" || os.Args[2] != "alpine" {
		fmt.Printf("Expected 'run alpine <command>' as the command! args: %q\n", os.Args[1:])
		os.Exit(1)
	}

	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Printf("TempDir Error: %v", err)
		os.Exit(1)
	}

	if err = os.Mkdir(tempDir+"/usr", os.ModeDir); err != nil {
		fmt.Printf("Mkdir error: %v", err)
		os.Exit(1)
	}

	if err = os.Mkdir(tempDir+"/usr/bin", os.ModeDir); err != nil {
		fmt.Printf("Mkdir error: %v", err)
		os.Exit(1)
	}

	if err := os.Symlink("/usr/bin/echo", tempDir+"/usr/bin/echo"); err != nil {
		fmt.Printf("Symlink Error: %v", err)
		os.Exit(1)
	}

	fmt.Println(tempDir)

	if err := syscall.Chroot(tempDir); err != nil {
		fmt.Printf("Chroot Error: %v", err)
		os.Exit(1)
	}

	if err := syscall.Exec("/usr/bin/"+os.Args[3], os.Args[3:], os.Environ()); err != nil {
		fmt.Printf("Exec Error: %v", err)
		os.Exit(1)
	}
}
