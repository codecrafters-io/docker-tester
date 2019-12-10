package main

import (
	"os"
)

func main() {
	os.Exit(RunCLI(os.Args[1:]))
}
