package main

import (
	"testing"

	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

func TestStart(t *testing.T) {
	err := NewExecutable("/blah").Start()
	assert.Assert(t, cmp.ErrorContains(err, "no such file"))
	assert.Assert(t, cmp.ErrorContains(err, "/blah"))

	err = NewExecutable("./test_helpers/executable_test/stdout_echo.sh").Start()
	assert.Assert(t, cmp.Nil(err))
}

func TestRun(t *testing.T) {
	e := NewExecutable("./test_helpers/executable_test/stdout_echo.sh")
	result, err := e.Run("hey")
	assert.Assert(t, cmp.Nil(err))
	assert.Assert(t, cmp.Equal(string(result.Stdout), "hey\n"))
}

func TestOutputCapture(t *testing.T) {
	// Stdout capture
	e := NewExecutable("./test_helpers/executable_test/stdout_echo.sh")
	result, err := e.Run("hey")

	assert.Assert(t, cmp.Nil(err))
	assert.Assert(t, cmp.Equal(string(result.Stdout), "hey\n"))
	assert.Assert(t, cmp.Equal(string(result.Stderr), ""))

	// Stderr capture
	e = NewExecutable("./test_helpers/executable_test/stderr_echo.sh")
	result, err = e.Run("hey")

	assert.Assert(t, cmp.Nil(err))
	assert.Assert(t, cmp.Equal(string(result.Stdout), ""))
	assert.Assert(t, cmp.Equal(string(result.Stderr), "hey\n"))
}

func TestExitCode(t *testing.T) {
	e := NewExecutable("./test_helpers/executable_test/exit_with.sh")

	result, _ := e.Run("0")
	assert.Assert(t, cmp.Equal(0, result.ExitCode))

	result, _ = e.Run("1")
	assert.Assert(t, cmp.Equal(1, result.ExitCode))
}

func TestExecutableStartNotAllowedIfInProgress(t *testing.T) {
	e := NewExecutable("./test_helpers/executable_test/sleep_for.sh")

	// Run once
	err := e.Start("0.01")
	assert.Assert(t, cmp.Nil(err))

	// Starting again when in progress should throw an error
	err = e.Start("0.01")
	assert.Assert(t, cmp.ErrorContains(err, "process already in progress"))

	// Running again when in progress should throw an error
	_, err = e.Run("0.01")
	assert.Assert(t, cmp.ErrorContains(err, "process already in progress"))

	e.Wait()

	// Running again once finished should be fine
	err = e.Start("0.01")
	assert.Assert(t, cmp.Nil(err))
}

func TestSuccessiveExecutions(t *testing.T) {
	e := NewExecutable("./test_helpers/executable_test/stdout_echo.sh")

	result, _ := e.Run("1")
	assert.Assert(t, cmp.Equal(string(result.Stdout), "1\n"))

	result, _ = e.Run("2")
	assert.Assert(t, cmp.Equal(string(result.Stdout), "2\n"))
}
