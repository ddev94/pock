#!/bin/bash

# Build script for pock

set -e

echo "Building pock..."

# Clean previous builds
rm -rf bin/
mkdir -p bin/

# Build for current platform
go build -o bin/pock ./cmd/pock

echo "Build completed successfully!"
echo "Binary location: bin/pock"
echo ""
echo "To install globally, run:"
echo "  sudo mv bin/pock /usr/local/bin/"
echo ""
echo "Or add bin/ to your PATH:"
echo "  export PATH=\$PATH:$(pwd)/bin"
