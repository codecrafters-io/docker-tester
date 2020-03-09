package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExec(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	fmt.Println("Test success")
	exitCode := runCLIStage("init", "./test_helpers/stages/basic_exec_success")
	if !assert.Equal(t, 0, exitCode) {
		failWithMockerOutput(t, m)
	}

	m.Reset()

	fmt.Println("Test failure")
	exitCode = runCLIStage("init", "./test_helpers/stages/basic_exec_failure")
	if !assert.Equal(t, 1, exitCode) {
		failWithMockerOutput(t, m)
	}
	assert.Contains(t, m.ReadStdout(), "Test failed")
}

func TestFSIsolation(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	// Previous solution should fail
	exitCode := runCLIStage("fs_isolation", "./test_helpers/stages/basic_exec_success")
	if !assert.Equal(t, 1, exitCode) {
		failWithMockerOutput(t, m)
	}

	m.Reset()

	// Current solution should succeed
	exitCode = runCLIStage("fs_isolation", "./test_helpers/stages/fs_isolation_success")
	if !assert.Equal(t, 0, exitCode) {
		failWithMockerOutput(t, m)
	}
}

// func TestProcessIsolation(t *testing.T) {
// 	m := NewStdIOMocker()
// 	m.Start()
// 	defer m.End()

// 	// Previous stage should fail
// 	exitCode := runCLIStage("./test_helpers/stages/fs_isolation/correct.sh", 2)
// 	if !assert.Equal(t, 1, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}

// 	m.Reset()

// 	// Next stage should succeed
// 	exitCode = runCLIStage("./test_helpers/stages/process_isolation/correct.sh", 2)
// 	if !assert.Equal(t, 0, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}
// }

// func TestFetchFromRegistry(t *testing.T) {
// 	m := NewStdIOMocker()
// 	m.Start()
// 	defer m.End()

// 	// Previous stage should fail
// 	exitCode := runCLIStage("./test_helpers/stages/process_isolation/correct.sh", 3)
// 	if !assert.Equal(t, 1, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}

// 	m.Reset()

// 	// Next stage should succeed
// 	exitCode = runCLIStage("./test_helpers/stages/fetch_from_registry/correct.sh", 3)
// 	if !assert.Equal(t, 0, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}
// }

func runCLIStage(slug string, path string) (exitCode int) {
	return RunCLI(map[string]string{
		"CODECRAFTERS_CURRENT_STAGE_SLUG": slug,
		"CODECRAFTERS_SUBMISSION_DIR":     path,
	})
}

func failWithMockerOutput(t *testing.T, m *IOMocker) {
	t.Error(fmt.Sprintf("stdout: \n%s\n\nstderr: \n%s", m.ReadStdout(), m.ReadStderr()))
}
