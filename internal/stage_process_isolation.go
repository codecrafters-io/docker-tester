package internal

import tester_utils "github.com/codecrafters-io/tester-utils"

func testProcessIsolation(stageHarness tester_utils.StageHarness) error {
	logger := stageHarness.Logger
	executable := stageHarness.Executable

	logger.Debugln("Running 'mypid'")
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
