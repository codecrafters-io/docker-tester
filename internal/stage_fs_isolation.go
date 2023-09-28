package internal

import (
	"os"
	"path/filepath"
	"strconv"

	tester_utils "github.com/codecrafters-io/tester-utils"
)

func testFSIsolation(stageHarness *tester_utils.StageHarness) error {
	initRandom()

	logger := stageHarness.Logger
	executable := stageHarness.Executable

	// tempDir, err := ioutil.TempDir("", "")
	tempDir, err := customTempDir()
	if err != nil {
		return err
	}

	logger.Debugf("Created temp dir on host: %v", tempDir)
	logger.Debugf("Executing: ./your_docker.sh run <some_image> /usr/local/bin/docker-explorer ls %v", tempDir)
	logger.Debugf("(expecting that the directory won't be accessible)")

	result, err := executable.Run(
		"run", DOCKER_IMAGE_STAGE_1,
		"/usr/local/bin/docker-explorer", "ls", tempDir,
	)
	if err != nil {
		return err
	}

	if err = assertStdoutContains(result, "No such file or directory"); err != nil {
		return err
	}

	if err = assertExitCode(result, 2); err != nil {
		return err
	}

	return nil
}

func customTempDir() (name string, err error) {
	dir := os.TempDir()

	dirname := strconv.Itoa(int(randomInt(99999)))
	fullPath := filepath.Join(dir, dirname)

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		if err = os.RemoveAll(fullPath); err != nil {
			return "", err
		}
	}

	err = os.Mkdir(fullPath, 0777)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}
