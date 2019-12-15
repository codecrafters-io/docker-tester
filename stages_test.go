package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExec(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	exitCode := runCLIStage("./test_helpers/stages/basic_exec/correct.sh", 0)
	if !assert.Equal(t, 0, exitCode) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	exitCode = runCLIStage("./test_helpers/stages/basic_exec/wrong.sh", 0)
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
	exitCode := runCLIStage("./test_helpers/stages/basic_exec/correct.sh", 1)
	if !assert.Equal(t, 1, exitCode) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	// Next stage should succeed
	exitCode = runCLIStage("./test_helpers/stages/fs_isolation/correct.sh", 1)
	if !assert.Equal(t, 0, exitCode) {
		t.Error(m.ReadStdout())
	}
}

func TestProcessIsolation(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	// Previous stage should fail
	exitCode := runCLIStage("./test_helpers/stages/fs_isolation/correct.sh", 2)
	if !assert.Equal(t, 1, exitCode) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	// Next stage should succeed
	exitCode = runCLIStage("./test_helpers/stages/fs_isolation/correct.sh", 2)
	if !assert.Equal(t, 0, exitCode) {
		t.Error(m.ReadStdout())
	}
}

func runCLIStage(path string, stage int) (exitCode int) {
	return RunCLI([]string{
		"--binary-path", path,
		"--stage", strconv.Itoa(stage),
		"--debug",
	})
}
