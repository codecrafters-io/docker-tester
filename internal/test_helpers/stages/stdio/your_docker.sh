#!/bin/sh
set -e
go build -o /tmp/out "$(dirname "$0")/main.go"
exec /tmp/out "$@"
