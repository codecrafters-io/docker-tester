package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestBasicExec(t *testing.T) {
	exitCode := RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec_correct.sh"})
	assert.Equal(t, exitCode, 1)
}
