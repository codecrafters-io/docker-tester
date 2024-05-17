package internal

import (
	"os"
	"testing"

	tester_utils_testing "github.com/codecrafters-io/tester-utils/testing"
)

func TestStages(t *testing.T) {
	os.Setenv("CODECRAFTERS_RANDOM_SEED", "1234567890")

	testCases := map[string]tester_utils_testing.TesterOutputTestCase{
		"basic_exec_failure": {
			UntilStageSlug:      "je9",
			CodePath:            "./test_helpers/stages/basic_exec_failure",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/basic_exec/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"basic_exec_success": {
			UntilStageSlug:      "je9",
			CodePath:            "./test_helpers/stages/basic_exec",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/basic_exec/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"stdio_failure": {
			UntilStageSlug:      "kf3",
			CodePath:            "./test_helpers/stages/basic_exec",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/stdio/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"stdio_success": {
			UntilStageSlug:      "kf3",
			CodePath:            "./test_helpers/stages/stdio",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/stdio/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"exit_code_success": {
			UntilStageSlug:      "kf3",
			CodePath:            "./test_helpers/stages/exit_code",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/exit_code/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"fs_isolation_failure": {
			UntilStageSlug:      "if6",
			CodePath:            "./test_helpers/stages/exit_code",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/fs_isolation/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"fs_isolation_success": {
			UntilStageSlug:      "if6",
			CodePath:            "./test_helpers/stages/fs_isolation",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/fs_isolation/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"process_isolation_success": {
			UntilStageSlug:      "lu7",
			CodePath:            "./test_helpers/stages/process_isolation",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/process_isolation/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"fetch_from_registry_failure": {
			UntilStageSlug:      "hs1",
			CodePath:            "./test_helpers/stages/process_isolation",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/fetch_from_registry/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"fetch_from_registry_success": {
			StageSlugs:          []string{"hs1"},
			CodePath:            "./test_helpers/stages/fetch_from_registry",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/fetch_from_registry/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
