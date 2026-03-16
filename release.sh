#!/bin/bash

# Release script for pock
set -e

# Get version from version.go
VERSION=$(grep 'const Version' pkg/pock/version.go | sed 's/.*"\(.*\)".*/\1/')

echo "========================================="
echo "Creating GitHub Release v${VERSION}"
echo "========================================="
echo ""

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "❌ GitHub CLI (gh) is not installed."
    echo ""
    echo "Install it with:"
    echo "  brew install gh"
    echo ""
    exit 1
fi

# Check if user is authenticated
if ! gh auth status &> /dev/null; then
    echo "❌ Not authenticated with GitHub."
    echo ""
    echo "Run: gh auth login"
    echo ""
    exit 1
fi

# Check if tag already exists
if git rev-parse "v${VERSION}" >/dev/null 2>&1; then
    echo "⚠️  Tag v${VERSION} already exists."
    read -p "Do you want to delete and recreate it? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        git tag -d "v${VERSION}"
        git push origin ":refs/tags/v${VERSION}" 2>/dev/null || true
    else
        echo "Aborted."
        exit 1
    fi
fi

# Build release artifacts
echo "▶ Building release artifacts..."
make release

echo ""
echo "▶ Creating git tag v${VERSION}..."
git tag -a "v${VERSION}" -m "Release v${VERSION}"

echo "▶ Pushing tag to GitHub..."
git push origin "v${VERSION}"

echo ""
echo "▶ Creating GitHub release..."
gh release create "v${VERSION}" \
    dist/pock-${VERSION}.pkg \
    dist/pock-${VERSION}-1.x86_64.tar.gz \
    dist/pock-${VERSION}-windows-amd64.zip \
    bin/pock-darwin-amd64#pock-darwin-amd64 \
    bin/pock-darwin-arm64#pock-darwin-arm64 \
    bin/pock-linux-amd64#pock-linux-amd64 \
    bin/pock-windows-amd64.exe#pock-windows-amd64.exe \
    --title "v${VERSION}" \
    --notes "Release v${VERSION}

## Installation

### macOS
Download and install \`pock-${VERSION}.pkg\`

**Universal binary** (works on both Intel and Apple Silicon Macs)

### Linux
Download \`pock-${VERSION}-1.x86_64.tar.gz\` and extract:
\`\`\`bash
tar -xzf pock-${VERSION}-1.x86_64.tar.gz
sudo cp pock-${VERSION}/usr/local/bin/pock /usr/local/bin/
\`\`\`

### Windows
Download and extract \`pock-${VERSION}-windows-amd64.zip\`
Add \`pock.exe\` to your PATH.

### Single Binaries
Download the binary for your platform:
- \`pock-darwin-amd64\` - macOS Intel
- \`pock-darwin-arm64\` - macOS Apple Silicon
- \`pock-linux-amd64\` - Linux 64-bit
- \`pock-windows-amd64.exe\` - Windows 64-bit

## What's Changed
- Bug fixes and improvements"

echo ""
echo "========================================="
echo "✓ Release v${VERSION} published!"
echo "========================================="
echo ""
echo "View at: $(gh release view "v${VERSION}" --json url -q .url)"
