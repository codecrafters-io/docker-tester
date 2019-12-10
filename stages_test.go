package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExec(t *testing.T) {
	m := NewStdIOMocker()
	m.Start()
	defer m.End()

	exitCode := RunCLI([]string{"--binary-path", "./test_helpers/stages/basic_exec_correct.sh"})
	assert.Equal(t, exitCode, 1)
}
