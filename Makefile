.PHONY: build run clean install test help package release

# Extract version from pkg/pock/version.go
VERSION := $(shell grep 'const Version' pkg/pock/version.go | sed 's/.*"\(.*\)".*/\1/')

# Run in development mode
run:
	@go run . $(ARGS)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean
	@echo "Clean completed!"

# Install globally
install: build
	@echo "Installing pock to /usr/local/bin..."
	@sudo mv bin/pock /usr/local/bin/
	@echo "Installation completed!"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded!"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting completed!"

# Lint code
lint:
	@echo "Linting code..."
	@golangci-lint run ./...
	@echo "Linting completed!"

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/pock-linux-amd64 .
	@GOOS=darwin GOARCH=amd64 go build -o bin/pock-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o bin/pock-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o bin/pock-windows-amd64.exe .
	@echo "Build completed for all platforms!"

release:
	@echo "========================================="
	@echo "Building release v$(VERSION) for all platforms..."
	@echo "========================================="
	@rm -rf bin
	
	@echo ""
	@echo "▶ Building cross-platform binaries..."
	@$(MAKE) -s build-all
	
	@echo ""
	@echo "========================================="
	@echo "✓ Release v$(VERSION) complete!"
	@echo "========================================="
	@echo ""
	@echo "Binaries ready in bin/:"
	@ls -lh bin/ 2>/dev/null || true
	@echo ""
	@echo "To create a GitHub release, run:"
	@echo "  ./release.sh"
