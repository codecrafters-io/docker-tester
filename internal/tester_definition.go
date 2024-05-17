package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{},
	ExecutableFileName: "your_docker.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "je9",
			TestFunc: testBasicExec,
		},
		{
			Slug:     "kf3",
			TestFunc: testStdio,
		},
		{
			Slug:     "cn8",
			TestFunc: testExitCode,
		},
		{
			Slug:     "if6",
			TestFunc: testFSIsolation,
		},
		{
			Slug:     "lu7",
			TestFunc: testProcessIsolation,
		},
		{
			Slug:     "hs1",
			TestFunc: testFetchFromRegistry,
		},
	},
}
