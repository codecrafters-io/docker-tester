package main

import "fmt"

func testBasicExec(executable *Executable, logger *customLogger) error {
	logger.Debugf("Executing 'echo foo'")
	result, err := executable.Run("run alpine echo foo")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "foo\n"); err != nil {
		return err
	}

	logger.Debugf("Executing 'echo bar'")
	result, err = executable.Run("run alpine echo bar")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "bar\n"); err != nil {
		return err
	}

	logger.Debugf("Executing 'exit 1'")
	result, err = executable.Run("run alpine exit 1")
	if err != nil {
		return err
	}

	if err = assertExitCode(result, 1); err != nil {
		return err
	}

	return nil
}

func assertStdout(result ExecutableResult, expected string) error {
	actual := string(result.Stdout)
	if expected != actual {
		return fmt.Errorf("Expected %q as stdout, got: %q", expected, actual)
	}

	return nil
}

func assertExitCode(result ExecutableResult, expected int) error {
	actual := result.ExitCode
	if expected != actual {
		return fmt.Errorf("Expected %d as exit code, got: %d", expected, actual)
	}

	return nil
}
