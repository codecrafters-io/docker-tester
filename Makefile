.PHONY: release build run

current_version_number := $(shell git tag --list "v*" | sort -V | tail -n 1 | cut -c 2-)
next_version_number := $(shell echo $$(($(current_version_number)+1)))

release:
	git tag v$(next_version_number)
	git push origin master v$(next_version_number)

build:
	go build -o dist/main.out ./cmd/tester

test:
	go test -v ./internal/

test_with_docker: build
	CODECRAFTERS_SUBMISSION_DIR=$(shell pwd)/internal/test_helpers/pass_all \
	CODECRAFTERS_TEST_CASES_JSON="[{"slug":"init","tester_log_prefix":"stage-1","title":"Stage #1: Execute a program"},{"slug":"stdio","tester_log_prefix":"stage-2","title":"Stage #2: Wireup stdout \u0026 stderr"},{"slug":"exit_code","tester_log_prefix":"stage-3","title":"Stage #3: Handle exit codes"},{"slug":"fs_isolation","tester_log_prefix":"stage-4","title":"Stage #4: Filesystem isolation"},{"slug":"process_isolation","tester_log_prefix":"stage-5","title":"Stage #5: Process isolation"},{"slug":"fetch_from_registry","tester_log_prefix":"stage-6","title":"Stage #6: Fetch an image from the Docker Registry"}]" \
	dist/main.out

test_in_docker_container:
	docker build -t docker-tester-dev . && docker run --cap-add "SYS_ADMIN" -e "TERM=xterm-256color" docker-tester-dev make test

test_in_docker_container_and_update_fixtures:
	docker build -t docker-tester-dev . && \
	docker run --name temp_test_container --cap-add "SYS_ADMIN" \
		-e "TERM=xterm-256color" \
		-e "CODECRAFTERS_RECORD_FIXTURES=true" \
		docker-tester-dev make test
	docker cp temp_test_container:/app/internal/test_helpers/fixtures ./internal/test_helpers/
	docker stop temp_test_container
	docker rm temp_test_container

copy_course_file:
	hub api \
		repos/rohitpaulk/codecrafters-server/contents/codecrafters/store/data/docker.yml \
		| jq -r .content \
		| base64 -d \
		> internal/test_helpers/course_definition.yml

test_output_failure_run:
	time sh -c "while true; do go test -run TestRun -v executable_test.go executable.go || break; done"

test_output_failure_start:
	time sh -c "while true; do go test -run TestStart -v executable_test.go executable.go || break; done"

test_output_failure:
	time sh -c "while true; do go test -v executable_test.go executable.go || break; done"

update_tester_utils:
	go get -u github.com/codecrafters-io/tester-utils