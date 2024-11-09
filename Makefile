# Makefile for idionautic-server

APP_NAME := idionautic-server
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
VERSION := $(shell git describe --tags --abbrev=0 || echo "v0.1.0")
BUILD_DIR := ./bin
DB_FILE := ./telemetry.db

.PHONY: all clean build run test lint fmt vet db-init db-drop

# Run everything: build, format, lint, and test
all: fmt lint test build

# Clean up built binaries and temporary files
clean:
	rm -rf $(BUILD_DIR) $(DB_FILE)

# Build the application binary if any Go files change
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): $(GO_FILES)
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) -ldflags "-X main.version=$(VERSION)" .

# Run the application (assumes build target has been run)
run: build
	$(BUILD_DIR)/$(APP_NAME)

# Run tests
test:
	go test ./...

# Run static analysis with Go vet
vet:
	go vet ./...

# Run code formatting
fmt:
	go fmt ./...

# Run linter (requires golangci-lint to be installed)
lint:
	golangci-lint run ./...

# Initialize the SQLite database with schema.sql (creates new DB if it doesn't exist)
db-init:
	@echo "Initializing database..."
	sqlite3 $(DB_FILE) < db/schema.sql || echo "Database already exists."

# Drop the SQLite database
db-drop:
	@echo "Dropping database..."
	rm -f $(DB_FILE)
