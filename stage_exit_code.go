package main

import (
	"math/rand"
	"strconv"
)

func testExitCode(executable *Executable, logger *customLogger) error {
	randomInt := rand.Intn(30)
	randomStr := strconv.Itoa(randomInt)

	logger.Debugf("Executing 'exit %s'", randomStr)
	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "exit", randomStr,
	)
	if err != nil {
		return err
	}

	logger.Debugf("Checking the parent process's exit code..")
	return assertExitCode(result, randomInt)
}
