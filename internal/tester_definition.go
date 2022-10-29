package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testerDefinition = testerutils.TesterDefinition{
	AntiCheatStages:    []testerutils.Stage{},
	ExecutableFileName: "your_docker.sh",
	Stages: []testerutils.Stage{
		{
			Number:                  1,
			Slug:                    "init",
			Title:                   "Execute a program",
			TestFunc:                testBasicExec,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  2,
			Slug:                    "stdio",
			Title:                   "Wireup stdout & stderr",
			TestFunc:                testStdio,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  3,
			Slug:                    "exit_code",
			Title:                   "Handle exit codes",
			TestFunc:                testExitCode,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  4,
			Slug:                    "fs_isolation",
			Title:                   "Filesystem isolation",
			TestFunc:                testFSIsolation,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  5,
			Slug:                    "process_isolation",
			Title:                   "Process isolation",
			TestFunc:                testProcessIsolation,
			ShouldRunPreviousStages: true,
		},
		{
			Number:                  6,
			Slug:                    "fetch_from_registry",
			Title:                   "Fetch an image from the Docker Registry",
			TestFunc:                testFetchFromRegistry,
			ShouldRunPreviousStages: false, // Takes too long!
		},
	},
}
