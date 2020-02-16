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

	exitCode := runCLIStage("./test_helpers/stages/basic_exec_success")
	if !assert.Equal(t, 0, exitCode) {
		failWithMockerOutput(t, m)
	}

	m.Reset()

	exitCode = runCLIStage("./test_helpers/stages/basic_exec_failure")
	if !assert.Equal(t, 1, exitCode) {
		failWithMockerOutput(t, m)
	}
	assert.Contains(t, m.ReadStdout(), "Test failed")
}

// func TestFSIsolation(t *testing.T) {
// 	m := NewStdIOMocker()
// 	m.Start()
// 	defer m.End()

// 	// Previous stage should fail
// 	exitCode := runCLIStage("./test_helpers/stages/basic_exec/correct.sh", 1)
// 	if !assert.Equal(t, 1, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}

// 	m.Reset()

// 	// Next stage should succeed
// 	exitCode = runCLIStage("./test_helpers/stages/fs_isolation/correct.sh", 1)
// 	if !assert.Equal(t, 0, exitCode) {
// 		failWithMockerOutput(t, m)
// 	}
// }

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

func runCLIStage(path string) (exitCode int) {
	return RunCLI(map[string]string{
		"APP_DIR": path,
	})
}

func failWithMockerOutput(t *testing.T, m *IOMocker) {
	t.Error(fmt.Sprintf("stdout: \n%s\n\nstderr: \n%s", m.ReadStdout(), m.ReadStderr()))
}
