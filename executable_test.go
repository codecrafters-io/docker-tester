package main

import (
	"testing"
)

func TestExecutableStart(t *testing.T) {
	e := NewExecutable("/blah")
	if err := e.Start(); err == nil {
		t.Error("Expected an error for an invalid executable path")
	}

	e = NewExecutable("/usr/bin/echo")
	if err := e.Start(); err != nil {
		t.Errorf("Did not expect an error for a valid executable path: %s", err)
	}
}

func TestExecutableWait(t *testing.T) {
	e := NewExecutable("/usr/bin/echo")
	if err := e.Start("hey"); err != nil {
		t.Errorf("Did not expect an error when running executable: %s", err)
	}

	result, err := e.Wait()
	if err != nil {
		t.Errorf("Did not expect an error when waiting on executable: %s", err)
	}
	if string(result.Stdout) != "hey\n" {
		t.Errorf("Expected hey, got: %q", result.Stdout)
	}
}
