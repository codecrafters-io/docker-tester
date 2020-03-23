.PHONY: release build run

current_version_number := $(shell git tag --list "v*" | sort -V | tail -n 1 | cut -c 2-)
next_version_number := $(shell echo $$(($(current_version_number)+1)))

release:
	git tag v$(next_version_number)
	git push origin master v$(next_version_number)

build:
	go build -o dist/main.out

test:
	go test -v

copy_course_file:
	hub api \
		repos/rohitpaulk/codecrafters-server/contents/codecrafters/store/data/docker.yml \
		| jq -r .content \
		| base64 -d \
		> test_helpers/course_definition.yml

test_output_failure:
	time sh -c "while true; do go test -v executable_test.go executable.go || break; done"
