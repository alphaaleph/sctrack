# build go server
FROM golang:1.20 as builder-go
WORKDIR /app

# auto refresh uses .air.toml
# RUN go install github.com/cosmtrek/air@latest

# copy dependencies list and dowload modules
COPY ../go.mod go.sum ./
RUN go mod download

# copy everything and build the app
COPY .. .
RUN go mod tidy
RUN go build -o ./server/bin ./...

# build archlinux image for the container
FROM archlinux:latest
WORKDIR /app

# copy everthing from the builder
#COPY --from=builder-react /app/build ./client
COPY --from=builder-go /app/server/bin/cmd .

# Copy the Swagger documentation files
COPY --from=builder-go /app/server/docs/ ./docs/

# expose the port and set the start
EXPOSE 3030
CMD ["./cmd"]
