FROM golang:1.20 as builder

# working directory
WORKDIR /app

# auto refresh uses .air.toml
# RUN go install github.com/cosmtrek/air@latest

# copy dependencies list and dowload modules
COPY go.mod go.sum ./
RUN go mod download

# copy everything and build the app
COPY . .
RUN go mod tidy
RUN go build -o ./bin ./...

# build archlinux image for the container
FROM archlinux:latest
WORKDIR /app

# copy everthing from the builder
COPY --from=builder /app/bin/cmd .

# expose the port and set the start
EXPOSE 3030
CMD ["./cmd"]
