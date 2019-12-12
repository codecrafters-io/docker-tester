package main

import (
	"io/ioutil"
)

func testFSIsolation(executable *Executable, logger *customLogger) error {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	logger.Debugf("Created temp dir on host: %v", tempDir)
	logger.Debugf("Executing 'ls' in %v", tempDir)

	result, err := executable.Run("run", "alpine", "ls", tempDir)
	if err != nil {
		return err
	}

	if err = assertStderrContains(result, "No such file or directory"); err != nil {
		return err
	}

	if err = assertExitCode(result, 2); err != nil {
		return err
	}

	return nil
}
