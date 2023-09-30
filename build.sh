#!/bin/bash

# Generate swagger
swag init --generatedTime --output ./server/docs --dir ./server/cmd/,./server/src/handlers/,./server/src/models

# build the docker images
echo "Build the docker image ..."
docker-compose build

# completed
echo "[Build Completed]"
