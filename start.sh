#!/bin/bash
set -e
cd "$(dirname "$0")"

export HTTP_PORT="8080"
export GIN_MODE=release
export REMONPI_VENDOR="mitsubishi"
export REMONPI_MODEL="kgsa3-c"
export REMONPI_DATABASE_PATH="$PWD/local"
export HEXPI_ADDRESS="http://localhost:8081"
go run ./cmd/remonpi/main.go
