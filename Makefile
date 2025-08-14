# Variables
APP_NAME=morse-to-text

# Tasks
default: build

#run:
#	@go run cmd/morse-to-text/main.go
build:
	@go build -o $(APP_NAME) cmd/morse-to-text/main.go
	@echo "Build completo: $(APP_NAME) gerado"
test:
	go test ./...
cover:
	go test -v -cover ./...