package internal

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
	logger.Debugf("(expecting that the directory won't be accessible)")

	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "ls", tempDir,
	)
	if err != nil {
		return err
	}

	if err = assertStdoutContains(result, "no such file or directory"); err != nil {
		return err
	}

	if err = assertExitCode(result, 2); err != nil {
		return err
	}

	return nil
}
