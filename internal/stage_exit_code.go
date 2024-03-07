package internal

import (
	"github.com/codecrafters-io/tester-utils/test_case_harness"
	"strconv"
)

func testExitCode(stageHarness *test_case_harness.TestCaseHarness) error {
	initRandom()

	logger := stageHarness.Logger
	executable := stageHarness.Executable

	randomStr := randomString(30)
	randomInt, _ := strconv.Atoi(randomStr)

	logger.Debugf("Executing: ./your_docker.sh run <some_image> /usr/local/bin/docker-explorer exit %s", randomStr)
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
