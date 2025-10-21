# Makefile for AWS HPC Platform
# Copyright 2025 Scott Friedman

.PHONY: help build test clean install lint fmt vet tidy

# Default target
help:
	@echo "AWS HPC Platform - Build Targets"
	@echo ""
	@echo "  make build    - Build CLI binary"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make install  - Install CLI to GOPATH/bin"
	@echo "  make lint     - Run linters"
	@echo "  make fmt      - Format code"
	@echo "  make vet      - Run go vet"
	@echo "  make tidy     - Tidy go modules"
	@echo ""

# Build variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -X 'github.com/aws-hpc/pkg.GitCommit=$(COMMIT)' \
           -X 'github.com/aws-hpc/pkg.BuildDate=$(BUILD_DATE)'

# Build CLI
build:
	@echo "Building aws-hpc CLI..."
	cd cli && go build -ldflags "$(LDFLAGS)" -o ../bin/aws-hpc .
	@echo "Binary: bin/aws-hpc"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./pkg/...
	go test -v ./cli/...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf dist/
	go clean

# Install CLI
install: build
	@echo "Installing aws-hpc to $(GOPATH)/bin..."
	cp bin/aws-hpc $(GOPATH)/bin/

# Run linters
lint:
	@echo "Running linters..."
	golangci-lint run ./...

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -w -s .
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Tidy modules
tidy:
	@echo "Tidying modules..."
	go mod tidy

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/aws-hpc-linux-amd64 ./cli
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/aws-hpc-linux-arm64 ./cli
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/aws-hpc-darwin-amd64 ./cli
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/aws-hpc-darwin-arm64 ./cli
	@echo "Binaries in dist/"
