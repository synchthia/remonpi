#!/bin/bash
set -e

export GIN_MODE=release
export REMONPI_VENDOR="mitsubishi"
export REMONPI_MODEL="kgsa3-c"
go run ./cmd/remonpi/main.go
