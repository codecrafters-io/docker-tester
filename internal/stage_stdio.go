package internal

import (
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testStdio(stageHarness *test_case_harness.TestCaseHarness) error {
	initRandom()

	logger := stageHarness.Logger
	executable := stageHarness.Executable

	randomStr := randomString(99999)

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

	randomStr = randomString(99999)
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
