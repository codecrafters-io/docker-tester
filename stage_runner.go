package main

import (
	"fmt"
	"math/rand"
	"time"
)

// StageRunnerResult is returned from StageRunner.Run()
type StageRunnerResult struct {
	lastStageIndex int
	error          error
	logger         *customLogger
}

// IsSuccess says whether a StageRunnerResult was successful
// or not
func (res StageRunnerResult) IsSuccess() bool {
	return res.error == nil
}

// StageRunner is used to run multiple stages
type StageRunner struct {
	stages  []Stage
	isDebug bool
}

func newStageRunner(isDebug bool) StageRunner {
	return StageRunner{
		isDebug: isDebug,
		stages: []Stage{
			Stage{
				slug:    "init",
				name:    "Stage 0: Execute a program",
				logger:  getLogger(isDebug, "[stage-0] "),
				runFunc: testBasicExec,
			},
			Stage{
				slug:    "stdio",
				name:    "Stage 1: Wireup stdout & stderr",
				logger:  getLogger(isDebug, "[stage-1] "),
				runFunc: testStdio,
			},
			Stage{
				slug:    "exit_code",
				name:    "Stage 1: Handle exit codes",
				logger:  getLogger(isDebug, "[stage-1] "),
				runFunc: testExitCode,
			},
			Stage{
				slug:    "fs_isolation",
				name:    "Stage 1: Filesystem isolation",
				logger:  getLogger(isDebug, "[stage-1] "),
				runFunc: testFSIsolation,
			},
			Stage{
				slug:    "process_isolation",
				name:    "Stage 2: Process isolation",
				logger:  getLogger(isDebug, "[stage-2] "),
				runFunc: testProcessIsolation,
			},
			Stage{
				slug:    "fetch_from_registry",
				name:    "Stage 3: Fetching images from a registry",
				logger:  getLogger(isDebug, "[stage-3] "),
				runFunc: testFetchFromRegistry,
			},
		},
	}
}

// Run tests in a specific StageRunner
func (r StageRunner) Run(executable *Executable) StageRunnerResult {
	for index, stage := range r.stages {
		logger := stage.logger
		logger.Infof("Running test: %s", stage.name)

		stageResultChannel := make(chan error, 1)
		go func() {
			err := stage.runFunc(executable, logger)
			stageResultChannel <- err
		}()

		var err error
		select {
		case stageErr := <-stageResultChannel:
			err = stageErr
		case <-time.After(120 * time.Second):
			err = fmt.Errorf("timed out, test exceeded 120 seconds")
		}

		if err != nil {
			reportTestError(err, r.isDebug, logger)
			return StageRunnerResult{
				lastStageIndex: index,
				error:          err,
			}
		}

		logger.Successf("Test passed.")
	}

	return StageRunnerResult{
		lastStageIndex: len(r.stages) - 1,
		error:          nil,
	}
}

func (r StageRunner) StageCount() int {
	return len(r.stages)
}

// Truncated returns a stageRunner with fewer stages
func (r StageRunner) Truncated(stageSlug string) StageRunner {
	newStages := make([]Stage, 0)
	for _, stage := range r.stages {
		newStages = append(newStages, stage)
		if stage.slug == stageSlug {
			return StageRunner{
				isDebug: r.isDebug,
				stages:  newStages,
			}
		}
	}

	panic(fmt.Sprintf("Stage slug %v not found. Stages: %v", stageSlug, r.stages))
}

// Fuck you, go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Randomized returns a stage runner that has stages randomized
func (r StageRunner) Randomized() StageRunner {
	return StageRunner{
		isDebug: r.isDebug,
		stages:  shuffleStages(r.stages),
	}
}

func shuffleStages(stages []Stage) []Stage {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Stage, len(stages))
	perm := r.Perm(len(stages))
	for i, randIndex := range perm {
		ret[i] = stages[randIndex]
	}
	return ret
}

func reportTestError(err error, isDebug bool, logger *customLogger) {
	logger.Errorf("%s", err)
	if isDebug {
		logger.Errorf("Test failed")
	} else {
		logger.Errorf("Test failed " +
			"(try setting 'debug: true' in your codecrafters.yml to see more details)")
	}
}

// Stage is blah
type Stage struct {
	slug    string
	name    string
	runFunc func(executable *Executable, logger *customLogger) error
	logger  *customLogger
}
