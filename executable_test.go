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

	err = NewExecutable("/usr/bin/echo").Start()
	assert.Assert(t, cmp.Nil(err))
}

func TestOutputCapture(t *testing.T) {
	// Stdout capture
	e := NewExecutable("./test_helpers/stdout_echo.sh")
	e.Start("hey")
	result, err := e.Wait()

	assert.Assert(t, cmp.Nil(err))
	assert.Assert(t, cmp.Equal(string(result.Stdout), "hey\n"))
	assert.Assert(t, cmp.Equal(string(result.Stderr), ""))

	// Stderr capture
	e = NewExecutable("./test_helpers/stderr_echo.sh")
	e.Start("hey")
	result, err = e.Wait()

	assert.Assert(t, cmp.Nil(err))
	assert.Assert(t, cmp.Equal(string(result.Stdout), ""))
	assert.Assert(t, cmp.Equal(string(result.Stderr), "hey\n"))
}

func TestExitCode(t *testing.T) {
	e := NewExecutable("./test_helpers/exit_with.sh")

	e.Start("0")
	result, _ := e.Wait()
	assert.Assert(t, cmp.Equal(0, result.ExitCode))

	e.Start("1")
	result, _ = e.Wait()
	assert.Assert(t, cmp.Equal(1, result.ExitCode))
}

func TestExecutableStartNotAllowedIfInProgress(t *testing.T) {
	e := NewExecutable("./test_helpers/sleep_for.sh")

	// Run once
	err := e.Start("0.01")
	assert.Assert(t, cmp.Nil(err))

	// Running again when in progress should throw an error
	err = e.Start("0.01")
	assert.Assert(t, cmp.ErrorContains(err, "process already in progress"))

	e.Wait()

	// Running again once finished should be fine
	err = e.Start("0.01")
	assert.Assert(t, cmp.Nil(err))
}

func TestSuccessiveExecutions(t *testing.T) {
	e := NewExecutable("./test_helpers/stdout_echo.sh")

	e.Start("1")
	result, _ := e.Wait()
	assert.Assert(t, cmp.Equal(string(result.Stdout), "1\n"))

	e.Start("2")
	result, _ = e.Wait()
	assert.Assert(t, cmp.Equal(string(result.Stdout), "2\n"))
}
