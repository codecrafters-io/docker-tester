package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{},
	ExecutableFileName: "your_docker.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "init",
			TestFunc: testBasicExec,
		},
		{
			Slug:     "stdio",
			TestFunc: testStdio,
		},
		{
			Slug:     "exit_code",
			TestFunc: testExitCode,
		},
		{
			Slug:     "fs_isolation",
			TestFunc: testFSIsolation,
		},
		{
			Slug:     "process_isolation",
			TestFunc: testProcessIsolation,
		},
		{
			Slug:     "fetch_from_registry",
			TestFunc: testFetchFromRegistry,
		},
	},
}
