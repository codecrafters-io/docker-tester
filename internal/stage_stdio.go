package internal

import (
	tester_utils "github.com/codecrafters-io/tester-utils"
	"math/rand"
	"strconv"
)

func testStdio(stageHarness *tester_utils.StageHarness) error {
	logger := stageHarness.Logger
	executable := stageHarness.Executable

	randomStr := strconv.FormatInt(rand.Int63n(99999), 10)

	logger.Debugf("Executing: ./your_docker.sh run <some_image> /usr/local/bin/docker-explorer echo %s", randomStr)
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
	logger.Debugf("Executing: ./your_docker.sh run <some_image> /usr/local/bin/docker-explorer echo_stderr %s", randomStr)
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
