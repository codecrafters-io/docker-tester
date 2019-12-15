package main

func testProcessIsolation(executable *Executable, logger *customLogger) error {
	logger.Debugln("Running 'mypid'")
	result, err := executable.Run("run", DOCKER_IMAGE_STAGE_1, "mypid")
	if err != nil {
		return err
	}

	if err := assertStdout(result, "1"); err != nil {
		return err
	}

	return nil
}
