#!/bin/bash
set -e
cd "$(dirname $0)"

export REACT_APP_API_HOST="localhost:8080"
export BROWSER="none"
npm run start
