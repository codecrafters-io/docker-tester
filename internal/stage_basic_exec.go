package internal

import (
	"math/rand"
	"strconv"
)

func testBasicExec(executable *Executable, logger *customLogger) error {
	randomStr := strconv.FormatInt(rand.Int63n(99999), 10)

	logger.Debugf("Executing 'echo %s'", randomStr)
	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "echo", randomStr,
	)
	if err != nil {
		return err
	}

	logger.Debugf("Checking if the command output was echo-ed..")
	return assertStdoutContains(result, randomStr)
}
