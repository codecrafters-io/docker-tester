package internal

import "github.com/codecrafters-io/tester-utils/test_case_harness"

func testProcessIsolation(stageHarness *test_case_harness.TestCaseHarness) error {
	logger := stageHarness.Logger
	executable := stageHarness.Executable

	logger.Debugln("Executing: ./your_docker.sh run <some_image> /usr/local/bin/docker-explorer mypid")
	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "mypid",
	)
	if err != nil {
		return err
	}

	if err := assertStdout(result, "1\n"); err != nil {
		return err
	}

	return nil
}
