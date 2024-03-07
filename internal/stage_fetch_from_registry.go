package internal

import (
	"github.com/codecrafters-io/tester-utils/test_case_harness"
	"strings"
)

func testFetchFromRegistry(stageHarness *test_case_harness.TestCaseHarness) error {
	logger := stageHarness.Logger
	executable := stageHarness.Executable

	logger.Debugln("Executing: ./your_docker.sh run alpine /bin/sh -c '/bin/ls /'")

	result, err := executable.Run(
		"run", "alpine",
		"/bin/sh", "-c", "/bin/ls /",
	)
	if err != nil {
		return err
	}

	if err := assertExitCode(result, 0); err != nil {
		return err
	}

	expectedStdout := strings.Join([]string{
		"bin",
		"dev",
		"etc",
		"home",
		"lib",
		"media",
		"mnt",
		"opt",
		"proc",
		"root",
		"run",
		"sbin",
		"srv",
		"sys",
		"tmp",
		"usr",
		"var",
	}, "\n")

	if err := assertStdout(result, expectedStdout+"\n"); err != nil {
		return err
	}

	return nil
}
