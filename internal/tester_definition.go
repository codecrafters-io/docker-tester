package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testerDefinition = testerutils.TesterDefinition{
	AntiCheatStages: []testerutils.Stage{},
	ExecutableFileName: "your_docker.sh",
	Stages: []testerutils.Stage{
		{
			Slug:     "init",
			Title:    "Execute a Program",
			TestFunc: testBasicExec,
			ShouldRunPreviousStages: true,
		},
		{
			Slug:     "stdio",
			Title:    "Wireup stdout & stderr",
			TestFunc: testStdio,
			ShouldRunPreviousStages: true,
		},
		{
			Slug:     "exit_code",
			Title:    "Handle exit codes",
			TestFunc: testExitCode,
			ShouldRunPreviousStages: true,
		},
		{
			Slug:     "fs_isolation",
			Title:    "Filesystem isolation",
			TestFunc: testFSIsolation,
			ShouldRunPreviousStages: true,
		},
		{
			Slug:     "process_isolation",
			Title:    "Process isolation",
			TestFunc: testProcessIsolation,
			ShouldRunPreviousStages: true,
		},
		{
			Slug:     "fetch_from_registry",
			Title:    "Fetching images from a registry",
			TestFunc: testFetchFromRegistry,
			ShouldRunPreviousStages: false, // Takes too long!
		},
	},
}
