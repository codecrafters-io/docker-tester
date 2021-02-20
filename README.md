# Docker Challenge Tester

This is a program that validates your progress on the Docker challenge.

# Requirements for binary

- Following environment variables:
  - `CODECRAFTERS_SUBMISSION_DIR` - root of the user's code submission
  - `CODECRAFTERS_CURRENT_STAGE_SLUG` - current stage the user is on

# User code requirements

- A binary named `your_docker.sh` that runs the user's docker command
- A file named `codecrafters.yml`, with the following values: 
  - `debug`

# Running tests

- `make test` on a Linux machine
- `make test_in_docker_container` on OSX
