#!/bin/bash

docker run -ti --rm \
    --name gotestmetrics \
    -p 8080:8080 \
    hervenicol/gotestmetrics
