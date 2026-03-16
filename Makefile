.PHONY: build run clean install test help package release

# Extract version from pkg/pock/version.go
VERSION := $(shell grep 'const Version' pkg/pock/version.go | sed 's/.*"\(.*\)".*/\1/')

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
	@GOOS=linux GOARCH=amd64 go build -o bin/pock-linux-amd64 .
	@GOOS=darwin GOARCH=amd64 go build -o bin/pock-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o bin/pock-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o bin/pock-windows-amd64.exe .
	@echo "Build completed for all platforms!"

# Build .deb package for Debian/Ubuntu (requires dpkg-deb, run on Linux or use Docker)
deb:
	@rm -rf dist/deb
	@mkdir -p dist/deb/pock_$(VERSION)/DEBIAN
	@mkdir -p dist/deb/pock_$(VERSION)/usr/local/bin
	
	@echo "Building Linux binary..."
	@GOOS=linux GOARCH=amd64 go build -o dist/deb/pock_$(VERSION)/usr/local/bin/pock .
	@chmod 755 dist/deb/pock_$(VERSION)/usr/local/bin/pock
	
	@echo "Package: pock" > dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Version: $(VERSION)" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Section: utils" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Priority: optional" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Architecture: amd64" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Maintainer: Pock Developers <dev@pock.io>" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo "Description: A simple app for saving and reusing terminal commands" >> dist/deb/pock_$(VERSION)/DEBIAN/control
	@echo " Keep your most-used commands in one place and run them anytime." >> dist/deb/pock_$(VERSION)/DEBIAN/control
	
	@if command -v dpkg-deb >/dev/null 2>&1; then \
		dpkg-deb --build dist/deb/pock_$(VERSION); \
		mv dist/deb/pock_$(VERSION).deb dist/pock_$(VERSION)_amd64.deb; \
		rm -rf dist/deb; \
		echo ""; \
		echo "✓ Debian package created: dist/pock_$(VERSION)_amd64.deb"; \
		echo ""; \
		echo "To install on Debian/Ubuntu:"; \
		echo "  sudo dpkg -i dist/pock_$(VERSION)_amd64.deb"; \
		echo ""; \
		echo "Or double-click the .deb file in the file manager"; \
	else \
		echo ""; \
		echo "⚠️  dpkg-deb not found. Package structure created but not built."; \
		echo ""; \
		echo "To build on a Linux machine with dpkg-deb:"; \
		echo "  dpkg-deb --build dist/deb/pock_$(VERSION)"; \
		echo ""; \
		echo "Or use Docker:"; \
		echo "  docker run --rm -v \$$(pwd):/work -w /work debian:latest sh -c 'apt-get update && apt-get install -y dpkg-dev && dpkg-deb --build dist/deb/pock_$(VERSION) && mv dist/deb/pock_$(VERSION).deb dist/pock_$(VERSION)_amd64.deb'"; \
	fi

# Build .rpm package for RedHat/Fedora/CentOS
rpm:
	@echo "Building tarball for RPM-based systems..."
	@rm -rf dist/rpm
	@mkdir -p dist/rpm/pock-$(VERSION)/usr/local/bin
	
	@GOOS=linux GOARCH=amd64 go build -o dist/rpm/pock-$(VERSION)/usr/local/bin/pock .
	@chmod 755 dist/rpm/pock-$(VERSION)/usr/local/bin/pock
	
	@cd dist/rpm && tar czf ../pock-$(VERSION)-1.x86_64.tar.gz pock-$(VERSION)
	@rm -rf dist/rpm
	@echo ""
	@echo "✓ Tarball created: dist/pock-$(VERSION)-1.x86_64.tar.gz"
	@echo ""
	@echo "To install on Linux:"
	@echo "  tar -xzf dist/pock-$(VERSION)-1.x86_64.tar.gz"
	@echo "  sudo cp pock-$(VERSION)/usr/local/bin/pock /usr/local/bin/"
	@echo "  sudo chmod 755 /usr/local/bin/pock"

# Build all Linux packages
linux-packages: deb rpm
	@echo "✓ All Linux packages built successfully!"

# Build Windows installer (creates a zip with the executable)
windows:
	@echo "Building Windows package..."
	@rm -rf dist/windows
	@mkdir -p dist/windows/pock-$(VERSION)-windows
	
	@GOOS=windows GOARCH=amd64 go build -o dist/windows/pock-$(VERSION)-windows/pock.exe .
	@echo "pock v$(VERSION)" > dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "Installation Instructions:" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "1. Copy pock.exe to a directory in your PATH" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "2. Or add this directory to your PATH environment variable" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	@echo "Usage: pock --help" >> dist/windows/pock-$(VERSION)-windows/VERSION.txt
	
	@cd dist/windows && zip -r ../pock-$(VERSION)-windows-amd64.zip pock-$(VERSION)-windows
	@rm -rf dist/windows
	@echo ""
	@echo "✓ Windows package created: dist/pock-$(VERSION)-windows-amd64.zip"

# Create GitHub release artifacts for all platforms
release:
	@echo "========================================="
	@echo "Building release v$(VERSION) for all platforms..."
	@echo "========================================="
	@rm -rf dist
	@mkdir -p dist
	
	@echo ""
	@echo "▶ Building macOS universal binary and installer..."
	@$(MAKE) -s app
	
	@echo ""
	@echo "▶ Building Linux tarball..."
	@$(MAKE) -s rpm
	
	@echo ""
	@echo "▶ Building Debian package structure..."
	@$(MAKE) -s deb
	
	@echo ""
	@echo "▶ Building Windows package..."
	@$(MAKE) -s windows
	
	@echo ""
	@echo "▶ Building cross-platform binaries..."
	@$(MAKE) -s build-all
	
	@echo ""
	@echo "========================================="
	@echo "✓ Release v$(VERSION) complete!"
	@echo "========================================="
	@echo ""
	@echo "Release artifacts:"
	@echo "  macOS:     dist/pock-$(VERSION).pkg"
	@echo "  Linux:     dist/pock-$(VERSION)-1.x86_64.tar.gz"
	@echo "  Linux deb: dist/deb/pock_$(VERSION)/ (requires dpkg-deb to finalize)"
	@echo "  Windows:   dist/pock-$(VERSION)-windows-amd64.zip"
	@echo ""
	@echo "Additional binaries in bin/:"
	@ls -lh bin/ 2>/dev/null || true
	@echo ""
	@echo "To create a GitHub release:"
	@echo "  1. git tag v$(VERSION)"
	@echo "  2. git push origin v$(VERSION)"
	@echo "  3. gh release create v$(VERSION) \\"
	@echo "       dist/pock-$(VERSION).pkg \\"
	@echo "       dist/pock-$(VERSION)-1.x86_64.tar.gz \\"
	@echo "       dist/pock-$(VERSION)-windows-amd64.zip \\"
	@echo "       --title 'v$(VERSION)' \\"
	@echo "       --notes 'Release v$(VERSION)'"

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
	@echo "  make app        - Build macOS universal binary and .pkg"
	@echo "  make deb        - Build Debian/Ubuntu .deb package"
	@echo "  make rpm        - Build RedHat/Fedora .rpm package"
	@echo "  make linux-packages - Build all Linux packages"
	@echo "  make windows    - Build Windows package"
	@echo "  make release    - Build release artifacts for all platforms"
	@echo "  make help       - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run ARGS='add hello \"echo Hello\"'"
	@echo "  make run ARGS='list --stats'"
	@echo "  make run ARGS='run hello'"

# Set to your Developer ID certificate name or leave empty for unsigned
SIGN_IDENTITY ?=

app:
	@rm -rf pkgroot
	@mkdir -p dist
	
	@echo "Building universal binary (arm64 + amd64)..."
	@GOOS=darwin GOARCH=arm64 go build -o dist/pock-arm64 .
	@GOOS=darwin GOARCH=amd64 go build -o dist/pock-amd64 .
	@lipo -create -output dist/pock dist/pock-arm64 dist/pock-amd64
	@rm dist/pock-arm64 dist/pock-amd64

ifdef SIGN_IDENTITY
	@echo "Signing binary with $(SIGN_IDENTITY)..."
	@codesign --force --sign "$(SIGN_IDENTITY)" \
	  --options runtime \
	  --timestamp \
	  dist/pock
endif

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

ifdef SIGN_IDENTITY
	@echo "Building signed package..."
	@COPYFILE_DISABLE=1 pkgbuild \
	  --root pkgroot \
	  --identifier com.azoom.pock \
	  --version $(VERSION) \
	  --install-location / \
	  --sign "$(SIGN_IDENTITY)" \
	  dist/pock-$(VERSION).pkg
	@echo "Package signed with $(SIGN_IDENTITY)"
else
	@echo "Building unsigned package (for testing only)..."
	@COPYFILE_DISABLE=1 pkgbuild \
	  --root pkgroot \
	  --identifier com.azoom.pock \
	  --version $(VERSION) \
	  --install-location / \
	  dist/pock-$(VERSION).pkg
	@echo ""
	@echo "⚠️  WARNING: Package is UNSIGNED and may not install on other devices!"
	@echo "To install on the target device, the user must:"
	@echo "  1. Right-click the .pkg file and select 'Open'"
	@echo "  2. Click 'Open' in the security dialog"
	@echo ""
	@echo "For distribution, sign with: make package SIGN_IDENTITY='Developer ID Installer: Your Name'"
endif