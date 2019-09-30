#!/bin/bash
set -e

cd $(dirname "$0")

export GIN_MODE=release
export REMONPI_VENDOR="mitsubishi"
export REMONPI_MODEL="kgsa3-c"
export REMONPI_DB_PATH="$PATH/remonpi.json"
go run ./cmd/remonpi/main.go
