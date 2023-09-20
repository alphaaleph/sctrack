#!/bin/bash

# build the application
echo "Build the ScTrack application ..."
go build -o ./bin ./...

# Generate swagger
swag init -g /src/handlers/carrier.go /src/handlers/journal.go /src/handlers/todos.go

# build the docker images
echo "Build the docker image ..."
docker-compose build server

# completed
echo "[Build Completed]"
