package main

func testBasicExec(executable *Executable, logger *customLogger) error {
	logger.Debugf("Executing 'echo foo'")
	result, err := executable.Run("run", DOCKER_IMAGE_STAGE_1, "/bin/echo", "foo")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "foo\n"); err != nil {
		return err
	}

	logger.Debugf("Executing 'echo bar'")
	result, err = executable.Run("run", DOCKER_IMAGE_STAGE_1, "/bin/echo", "bar")
	if err != nil {
		return err
	}

	if err = assertStdout(result, "bar\n"); err != nil {
		return err
	}

	return nil
}
