.PHONY: build run clean install test help package

# Build the application
build:
	@echo "Building pock..."
	@mkdir -p bin
	@go build -o bin/pock ./cmd/pock
	@echo "Build completed! Binary: bin/pock"

# Run in development mode
run:
	@go run ./cmd/pock $(ARGS)

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
	@GOOS=linux GOARCH=amd64 go build -o bin/pock-linux-amd64 ./cmd/pock
	@GOOS=darwin GOARCH=amd64 go build -o bin/pock-darwin-amd64 ./cmd/pock
	@GOOS=darwin GOARCH=arm64 go build -o bin/pock-darwin-arm64 ./cmd/pock
	@GOOS=windows GOARCH=amd64 go build -o bin/pock-windows-amd64.exe ./cmd/pock
	@echo "Build completed for all platforms!"

# Show help
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run in development mode (use ARGS='...' for arguments)"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make install    - Install globally to /usr/local/bin"
	@echo "  make test       - Run tests"
	@echo "  make deps       - Download dependencies"
	@echo "  make fmt        - Format code"
	@echo "  make lint       - Lint code"
	@echo "  make build-all  - Build for all platforms"
	@echo "  make package    - Build macOS .pkg installer"
	@echo "  make help       - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run ARGS='add hello \"echo Hello\"'"
	@echo "  make run ARGS='list --stats'"
	@echo "  make run ARGS='run hello'"

package:
	@rm -rf pkgroot
	@mkdir -p dist
	@go build -o dist/pock ./cmd/pock

	@mkdir -p pkgroot/usr/local/bin
	@cp dist/pock pkgroot/usr/local/bin/pock
	@chmod 755 pkgroot/usr/local/bin/pock

	@mkdir -p pkgroot/usr/local/share/zsh/site-functions
	@mkdir -p pkgroot/usr/local/etc/bash_completion.d
	@mkdir -p pkgroot/usr/local/share/fish/vendor_completions.d

	@dist/pock completion zsh > pkgroot/usr/local/share/zsh/site-functions/_pock
	@dist/pock completion bash > pkgroot/usr/local/etc/bash_completion.d/pock
	@dist/pock completion fish > pkgroot/usr/local/share/fish/vendor_completions.d/pock.fish
	@chmod 644 pkgroot/usr/local/share/zsh/site-functions/_pock
	@chmod 644 pkgroot/usr/local/etc/bash_completion.d/pock
	@chmod 644 pkgroot/usr/local/share/fish/vendor_completions.d/pock.fish
	@find pkgroot -name '._*' -delete
	@xattr -cr pkgroot

	@COPYFILE_DISABLE=1 pkgbuild \
	  --root pkgroot \
	  --identifier com.azoom.pock \
	  --version 1.0.0 \
	  --install-location / \
	  dist/pock-1.0.0.pkg