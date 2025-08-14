# Variables
APP_NAME=morse-to-text

# Tasks
default: build

#run:
#	@go run cmd/morse-to-text/main.go
build:
	@go build -o $(APP_NAME) cmd/morse-to-text/main.go
test:
	go test ./...