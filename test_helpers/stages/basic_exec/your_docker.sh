#!/usr/bin/bash
set -e
go build -o /tmp/out "$(dirname "${BASH_SOURCE[0]}")/main.go" 
exec /tmp/out "$@"
