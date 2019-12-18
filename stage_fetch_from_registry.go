package main

import "strings"

func testFetchFromRegistry(executable *Executable, logger *customLogger) error {
	logger.Debugln("Running 'ls' using an alpine image")

	result, err := executable.Run("run", "alpine:3.10.3", "/bin/sh", "-c", "ls")
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

	if err := assertStdout(result, expectedStdout); err != nil {
		return err
	}

	return nil
}
