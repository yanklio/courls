# File targets - these represent actual files created by actions
BINARY := courls
BUILD_DIR := build
COVER_PROFILE := $(BUILD_DIR)/cover.out
COVER_HTML := $(BUILD_DIR)/cover.html

# Declare phony targets
.PHONY: run test test-verbose test-coverage clean fmt lint deps start

# Build the application - creates the binary file
$(BINARY): *.go go.mod go.sum
	go build -o $(BUILD_DIR)/$(BINARY) .

# Alias for the binary target
build: $(BINARY)

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Generate coverage profile
$(COVER_PROFILE): $(BUILD_DIR) *.go go.mod go.sum
	go test ./... -v --coverprofile $(COVER_PROFILE)

# Generate coverage HTML report
$(COVER_HTML): $(COVER_PROFILE)
	go tool cover -html $(COVER_PROFILE) -o $(COVER_HTML)

# Test the application
test:
	go test ./...

# Test with verbose output
test-verbose:
	go test -v ./...

# Test with coverage and generate HTML report in build directory
test-coverage: $(COVER_HTML)

# Clean build artifacts - removes tracked files
clean:
	rm -rf $(BUILD_DIR)

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download
	go mod tidy
