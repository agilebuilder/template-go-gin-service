APP_NAME := <%= projectName %>
BUILD_DIR := bin

.PHONY: build run test vet tidy clean

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server

run:
	go run ./cmd/server -config config.yaml

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BUILD_DIR)
