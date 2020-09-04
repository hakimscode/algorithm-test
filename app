#!/usr/bin/env bash
docker run -it --rm --name algorithm_test -v "$PWD":/app -w /app golang:1.14.4 go run . ${@}
