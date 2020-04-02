package internal

import (
	"math/rand"
	"strconv"
)

func testStdio(executable *Executable, logger *customLogger) error {
	randomStr := strconv.FormatInt(rand.Int63n(99999), 10)

	logger.Debugf("Executing 'echo %s'", randomStr)
	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "echo", randomStr,
	)
	if err != nil {
		return err
	}

	logger.Debugf("Checking if the string was echo-ed to stdout..")
	if err = assertStdout(result, randomStr+"\n"); err != nil {
		return err
	}

	randomStr = strconv.FormatInt(rand.Int63n(99999), 10)
	logger.Debugf("Executing 'echo_stderr %s'", randomStr)
	result, err = executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "echo_stderr", randomStr,
	)
	if err != nil {
		return err
	}

	logger.Debugf("Checking if the string was echo-ed to stderr..")
	if err = assertStderr(result, randomStr+"\n"); err != nil {
		return err
	}

	return nil
}
