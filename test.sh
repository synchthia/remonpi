#!/bin/bash
set -e
cd "$(dirname "$0")"

if [ "$1" = "remote" ]; then
    if [ "$2" = "get" ]; then
        set -x

        curl -fsSL \
            -X GET \
            -H 'Content-Type: application/json' \
            http://localhost:8080/api/v1/remote | jq .

        set +x
    fi
    if [ "$2" = "post" ]; then
        if [ "$3" = "on" ]; then
            set -x
            curl \
                -X POST \
                -H 'Content-Type: application/json' \
                -d '{"operation": true, "mode": "cool", "temp": 28, "fan": "high", "vertical_vane": "auto", "horizontal_vane": "keep"}' \
                http://localhost:8080/api/v1/remote | jq .
            set +x
        fi
        if [ "$3" = "full" ]; then
            set -x
            curl \
                -X POST \
                -H 'Content-Type: application/json' \
                -d '{"operation": true, "mode": "cool", "mode_data": {"cool": {"temp": 28, "fan": "auto", "vertical_vane": "auto", "horizontal_vane": "keep"}}}' \
                http://localhost:8080/api/v1/remote | jq .
            set +x
        fi
        if [ "$3" = "off" ]; then
            set -x
            curl \
                -X POST \
                -H 'Content-Type: application/json' \
                -d '{"operation": false, "mode": "cool", "temp": 28, "fan": "auto", "vertical_vane": "auto", "horizontal_vane": "keep"}' \
                http://localhost:8080/api/v1/remote | jq .
            set +x
        fi
        if [ "$3" = "dry" ]; then
            set -x
            curl \
                -X POST \
                -H 'Content-Type: application/json' \
                -d '{"operation": true, "mode": "dry", "fan": "auto", "vertical_vane": "auto", "horizontal_vane": "keep"}' \
                http://localhost:8080/api/v1/remote | jq .
            set +x
        fi

        if [ "$3" = "invalid" ]; then
            set -x
            curl \
                -X POST \
                -H 'Content-Type: application/json' \
                -d '{"operation": false, "mode": "cool", "temp": 28, "fan": "auto"}' \
                http://localhost:8080/api/v1/remote | jq .
            set +x
        fi

    fi
fi
