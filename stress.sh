#!/bin/bash

docker run -ti --rm \
    -v "$(pwd)"/stress:/home/k6/tests \
    --network host \
    grafana/k6 \
    run --vus 2 --duration 30m tests/script.js
