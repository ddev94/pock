#!/bin/bash

# Run script for development

set -e

echo "Running pock in development mode..."
go run ./cmd/pock "$@"
