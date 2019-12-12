package main

func testBasicExec(executable *Executable, logger *customLogger) error {
	logger.Debugf("Executing 'echo foo'")
	result, err := executable.Run("run", "alpine", "echo", "foo")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "foo\n"); err != nil {
		return err
	}

	logger.Debugf("Executing 'echo bar'")
	result, err = executable.Run("run", "alpine", "echo", "bar")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "bar\n"); err != nil {
		return err
	}

	logger.Debugf("Executing 'exit 1'")
	result, err = executable.Run("run", "alpine", "sh", "-c", "exit 1")
	if err != nil {
		return err
	}

	if err = assertExitCode(result, 1); err != nil {
		return err
	}

	logger.Debugf("Executing 'exit 2'")
	result, err = executable.Run("run", "alpine", "sh", "-c", "exit 2")
	if err != nil {
		return err
	}

	if err = assertExitCode(result, 2); err != nil {
		return err
	}

	return nil
}
