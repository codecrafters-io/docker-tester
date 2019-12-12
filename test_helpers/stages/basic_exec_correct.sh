#!/usr/bin/bash
go build -o /tmp/out "$(dirname "${BASH_SOURCE[0]}")/basic_exec_correct.go" 
/tmp/out "$@"
