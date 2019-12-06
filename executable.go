package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os/exec"
)

// Executable represents a program that can be executed
type Executable struct {
	path          string
	timeoutInSecs int

	// These are set & removed together
	cmd        *exec.Cmd
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
}

// ExecutableResult holds the result of an executable run
type ExecutableResult struct {
	Stdout   []byte
	Stderr   []byte
	ExitCode int
}

// NewExecutable returns an Executable struct
func NewExecutable(path string) *Executable {
	return &Executable{path: path, timeoutInSecs: 10}
}

func (e *Executable) isRunning() bool {
	return e.cmd != nil
}

// Start starts the specified command but does not wait for it to complete.
func (e *Executable) Start(args ...string) error {
	var err error

	if e.isRunning() {
		return errors.New("process already in progress")
	}

	// TODO: Use timeout!
	e.cmd = exec.Command(e.path, args...)

	e.stdoutPipe, err = e.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	e.stderrPipe, err = e.cmd.StderrPipe()
	if err != nil {
		return err
	}

	return e.cmd.Start()
}

// Wait waits for the program to finish and results the result
func (e *Executable) Wait() (ExecutableResult, error) {
	stdout, stdoutErr := ioutil.ReadAll(e.stdoutPipe)
	if stdoutErr != nil {
		return ExecutableResult{}, stdoutErr
	}
	stderr, stderrErr := ioutil.ReadAll(e.stderrPipe)
	if stderrErr != nil {
		return ExecutableResult{}, stderrErr
	}

	e.cmd.Wait()

	defer func() {
		e.cmd = nil
		e.stdoutPipe = nil
		e.stderrPipe = nil
	}()

	return ExecutableResult{
		Stdout:   stdout,
		Stderr:   stderr,
		ExitCode: e.cmd.ProcessState.ExitCode(),
	}, nil
}
