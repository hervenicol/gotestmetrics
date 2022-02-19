#!/bin/bash

docker run -ti --rm \
    -v "$(pwd)":/go/src/gotestmetrics \
    -w /go/src/gotestmetrics \
    -e CGO_ENABLED=0 \
    -e GOOS=linux \
    -e GOARCH=amd64 \
    golang:1.17 \
    go build -o gotestmetrics

docker build -t hervenicol/gotestmetrics .

