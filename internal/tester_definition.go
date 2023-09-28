package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testerDefinition = testerutils.TesterDefinition{
	AntiCheatTestCases:    []testerutils.TestCase{},
	ExecutableFileName: "your_docker.sh",
	TestCases: []testerutils.TestCase{
		{
			Slug:                    "init",
			TestFunc:                testBasicExec,
		},
		{
			Slug:                    "stdio",
			TestFunc:                testStdio,
		},
		{
			Slug:                    "exit_code",
			TestFunc:                testExitCode,
		},
		{
			Slug:                    "fs_isolation",
			TestFunc:                testFSIsolation,
		},
		{
			Slug:                    "process_isolation",
			TestFunc:                testProcessIsolation,
		},
		{
			Slug:                    "fetch_from_registry",
			TestFunc:                testFetchFromRegistry,
		},
	},
}
