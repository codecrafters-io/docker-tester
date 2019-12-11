package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExec(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	exitCode := RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec_correct.sh", "--stage", "0"})
	if !assert.Equal(t, exitCode, 0) {
		t.Error(m.ReadStdout())
	}

	m.Reset()

	exitCode = RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec_wrong.sh", "--stage", "0"})
	if !assert.Equal(t, exitCode, 1) {
		t.Error(m.ReadStdout())
	}
	assert.Contains(t, m.ReadStdout(), "Test failed")
}
}
