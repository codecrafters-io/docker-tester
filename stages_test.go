package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExec(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	exitCode := RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec/correct.sh", "--stage", "0"})
	if !assert.Equal(t, 0, exitCode) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	exitCode = RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec/wrong.sh", "--stage", "0"})
	if !assert.Equal(t, 1, exitCode) {
		t.Error(m.ReadStdout())
	}
	assert.Contains(t, m.ReadStdout(), "Test failed")
}

func TestFSIsolation(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	// Previous stage should fail
	exitCode := RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec/correct.sh", "--stage", "1"})
	if !assert.Equal(t, 1, exitCode) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	// Next stage should succeed
	exitCode = RunCLI([]string{"--binary-path", "./test_helpers/stages/fs_isolation/correct.sh", "--stage", "1"})
	if !assert.Equal(t, 0, exitCode) {
		t.Error(m.ReadStdout())
	}
}
