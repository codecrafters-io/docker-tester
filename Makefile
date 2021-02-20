.PHONY: release build run

current_version_number := $(shell git tag --list "v*" | sort -V | tail -n 1 | cut -c 2-)
next_version_number := $(shell echo $$(($(current_version_number)+1)))

release:
	git tag v$(next_version_number)
	git push origin master v$(next_version_number)

build:
	go build -o dist/main.out ./cmd/tester
	go build -o dist/starter.out ./cmd/starter_tester

test:
	go test -v ./internal/

test_in_docker_container:
	docker build -t docker-tester-dev . && docker run --cap-add "SYS_ADMIN" -it -e "TERM=xterm-256color" docker-tester-dev make test

copy_course_file:
	hub api \
		repos/rohitpaulk/codecrafters-server/contents/codecrafters/store/data/docker.yml \
		| jq -r .content \
		| base64 -d \
		> test_helpers/course_definition.yml

test_output_failure_run:
	time sh -c "while true; do go test -run TestRun -v executable_test.go executable.go || break; done"

test_output_failure_start:
	time sh -c "while true; do go test -run TestStart -v executable_test.go executable.go || break; done"

test_output_failure:
	time sh -c "while true; do go test -v executable_test.go executable.go || break; done"
