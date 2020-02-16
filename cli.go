package main

import (
	"fmt"
	"os"
	"os/exec"

	"syscall"
	"time"
)

// RunCLI executes the CLI program with given flags, and returns the exit code
func RunCLI(envMap map[string]string) int {
	fmt.Println("Welcome to the docker challenge!")
	fmt.Println("")

	context, err := GetContext(envMap)
	if err != nil {
		fmt.Printf("%s", err)
		return 1
	}

	context.print()
	fmt.Println("")

	executable := NewExecutable(context.binaryPath)
	if context.isDebug {
		executable = NewVerboseExecutable(context.binaryPath)
	}

	// TODO: Signal handlers!
	// installSignalHandler(cmd)

	runner := newStageRunner(context.isDebug)
	runner = runner.Truncated(context.currentStageIndex)

	_, err = runInOrder(runner, executable)
	if err != nil {
		return 1
	}

	if context.currentStageIndex > 0 {
		err = runRandomizedMultipleAndLog(runner, executable)
		if err != nil {
			return 1
		}
	}

	if antiCheatRunner().Run(executable).error != nil {
		return 1
	}

	return 0
}

func runRandomizedMultipleAndLog(runner StageRunner, executable *Executable) error {
	fmt.Println("Running tests multiple times to make sure...")

	fmt.Println("")
	time.Sleep(1 * time.Second)

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d...\n\n", i)
		time.Sleep(1 * time.Second)
		err := runRandomized(runner, executable)
		if err != nil {
			return err
		}
		fmt.Println("")
	}

	return nil
}

func runInOrder(runner StageRunner, executable *Executable) (StageRunnerResult, error) {
	result := runner.Run(executable)
	if !result.IsSuccess() {
		return result, fmt.Errorf("error")
	}

	fmt.Println("")
	fmt.Println("All tests ran successfully. Congrats!")
	fmt.Println("")
	return result, nil
}

func runRandomized(runner StageRunner, executable *Executable) error {
	result := runner.Randomized().Run(executable)
	if !result.IsSuccess() {
		return fmt.Errorf("error")
	}

	return nil
}

func runBinary(binaryPath string, debug bool) (*exec.Cmd, error) {
	command := exec.Command(binaryPath)
	if debug {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	err := command.Start()
	if err != nil {
		return nil, err
	}

	return command, nil
}
