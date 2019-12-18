# Go version >= v1.11 for support go mod
FROM golang:1.12

# Enable go mod support
ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080

# Run our Go built application name: demo-golang-app
# If you want to change the name of binary file when using go build
# run this command: go build -o myapp main.go
# Also change entrypoint from demo-golang-app -> myapp like this  
# ENTRYPOINT ["/app/myapp"]

ENTRYPOINT ["/app/demo-golang-app"]
