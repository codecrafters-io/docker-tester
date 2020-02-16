#!/usr/bin/bash
go build -o /tmp/out "$(dirname "${BASH_SOURCE[0]}")/main.go" 
/tmp/out "$@"
