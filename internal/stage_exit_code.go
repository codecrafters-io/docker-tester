package internal

import (
	tester_utils "github.com/codecrafters-io/tester-utils"
	"math/rand"
	"strconv"
)

func testExitCode(stageHarness *tester_utils.StageHarness) error {
	logger := stageHarness.Logger
	executable := stageHarness.Executable

	randomInt := rand.Intn(30)
	randomStr := strconv.Itoa(randomInt)

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
