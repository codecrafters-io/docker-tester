package main

import (
	"flag"
	"fmt"
)

// Context holds all flags that a user has passed in
type Context struct {
	binaryPath        string
	isDebug           bool
	currentStageIndex int
	reportOnSuccess   bool
	apiKey            string
}

func (c Context) print() {
	fmt.Println("Binary Path =", c.binaryPath)
	fmt.Println("Debug =", c.isDebug)
	fmt.Println("Report On Success =", c.reportOnSuccess)
	fmt.Println("Stage =", c.currentStageIndex)
}

var binaryPathPtr = flag.String(
	"binary-path",
	"",
	"path to the executable to test. Ex: ./run_program.sh")

var debugPtr = flag.Bool(
	"debug",
	false,
	"Whether debug logs must be printed")

var apiKeyPtr = flag.String(
	"api-key",
	"",
	"API key to use for reporting test results")

var reportOnSuccessPtr = flag.Bool(
	"report",
	false,
	"Whether test results must be reported")

var currentStagePtr = flag.Int(
	"stage",
	0,
	"The current stage you're on")

// GetContext parses flags and returns a Context object
func GetContext(args []string) (Context, error) {
	flag.CommandLine.Parse(args)

	if *binaryPathPtr == "" {
		return Context{}, fmt.Errorf("" +
			"The --binary-path flag must be specified")
	}

	if *reportOnSuccessPtr && (*apiKeyPtr == "") {
		return Context{}, fmt.Errorf("" +
			"If --report is specified, " +
			"--api-key must be specified too.")
	}

	return Context{
		binaryPath:        *binaryPathPtr,
		isDebug:           *debugPtr,
		currentStageIndex: *currentStagePtr,
		reportOnSuccess:   *reportOnSuccessPtr,
		apiKey:            *apiKeyPtr,
	}, nil
}
