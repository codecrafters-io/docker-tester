package internal

import (
	"fmt"
	testerutils "github.com/codecrafters-io/tester-utils"
)

func RunCLI(env map[string]string) int {
	tester, err := testerutils.NewTester(env, testerDefinition)

	if err != nil {
		fmt.Printf("%s", err)
		return 1
	}

	tester.PrintDebugContext()

	if !tester.RunStages() {
		tester.PrintFailureMessage()
		return 1
	}

	if !tester.RunAntiCheatStages() {
		return 1
	}

	tester.PrintSuccessMessage()
	return 0
}
