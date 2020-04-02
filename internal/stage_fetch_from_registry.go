package main

import "strings"

func testFetchFromRegistry(executable *Executable, logger *customLogger) error {
	logger.Debugln("Running 'ls' using an alpine image")

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
