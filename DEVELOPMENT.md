# Development Guide

## Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)

## Installation

### Installing Go

#### macOS

```bash
brew install go
```

#### Linux

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install golang

# Fedora
sudo dnf install golang

# Arch
sudo pacman -S go
```

#### Windows

Download and install from [golang.org](https://golang.org/dl/)

## Project Structure

```
golang/
├── cmd/
│   └── pock/
│       └── main.go              # Application entry point
├── internal/
│   ├── commands/                # Command implementations
│   │   ├── add.go
│   │   ├── browse.go
│   │   ├── config.go
│   │   ├── export.go
│   │   ├── history.go
│   │   ├── import.go
│   │   ├── install.go
│   │   ├── list.go
│   │   ├── publish.go
│   │   ├── remove.go
│   │   └── run.go
│   ├── storage/                 # Data storage layer
│   │   ├── commands.go          # Command CRUD operations
│   │   ├── database.go          # Database management
│   │   ├── history.go           # History operations
│   │   ├── settings.go          # Settings operations
│   │   └── types.go             # Type definitions
│   └── utils/                   # Utility functions
│       ├── colors.go            # Color formatting
│       ├── exec.go              # Command execution
│       └── table.go             # Table rendering
├── go.mod                       # Go module definition
├── Makefile                     # Build automation
├── build.sh                     # Build script
├── run.sh                       # Development run script
└── README.md                    # Main documentation
```

## Getting Started

### 1. Download Dependencies

```bash
cd golang
go mod download
```

### 2. Build the Application

#### Using Make

```bash
make build
```

#### Using build script

```bash
./build.sh
```

#### Using go directly

```bash
go build -o bin/pock ./cmd/pock
```

### 3. Run in Development Mode

#### Using Make

```bash
make run ARGS='list'
make run ARGS='add test "echo test"'
```

#### Using run script

```bash
./run.sh list
./run.sh add test "echo test"
```

#### Using go directly

```bash
go run ./cmd/pock list
go run ./cmd/pock add test "echo test"
```

## Building

### Build for Current Platform

```bash
make build
# or
go build -o bin/pock ./cmd/pock
```

### Build for All Platforms

```bash
make build-all
```

This creates binaries for:

- Linux (amd64)
- macOS (amd64 and arm64)
- Windows (amd64)

## Installation

### Local Installation

```bash
make install
# or
sudo mv bin/pock /usr/local/bin/
```

### Add to PATH (without sudo)

```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export PATH=$PATH:/path/to/command-line-manager/golang/bin
```

## Testing

### Run Tests

```bash
make test
# or
go test -v ./...
```

### Run Specific Tests

```bash
go test -v ./internal/storage/...
go test -v ./internal/commands/...
```

## Code Quality

### Format Code

```bash
make fmt
# or
go fmt ./...
```

### Lint Code

```bash
make lint
# or
golangci-lint run ./...
```

Note: You need to install golangci-lint first:

```bash
# macOS
brew install golangci-lint

# Linux
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

## Usage Examples

### Adding Commands

```bash
# Add a simple command
pock add hello "echo 'Hello, World!'"

# Add with description
pock add deploy "git push origin main" -d "Deploy to production"
```

### Listing Commands

```bash
# List all commands
pock list

# List with execution statistics
pock list --stats
```

### Running Commands

```bash
# Run a saved command
pock run hello
```

### Managing History

```bash
# View execution history
pock history

# View last 10 entries
pock history --limit 10

# Clear history
pock history --clear
```

### Exporting and Importing

```bash
# Export all commands
pock export commands.json

# Export specific command
pock export commands.json --name hello

# Import from file
pock import commands.json

# Import from URL
pock import https://example.com/commands.json

# Force overwrite existing commands
pock import commands.json --force
```

### Configuration

```bash
# List all settings
pock config list

# Get a setting value
pock config get listLayout

# Set a setting value
pock config set listLayout simple
pock config set dateFormat iso

# Reset to defaults
pock config reset
```

## Data Storage

The application stores data in JSON format at:

- macOS/Linux: `~/.local/share/pock/db.json`
- Windows: `%LOCALAPPDATA%\pock\db.json`

The database contains:

- Saved commands
- Command execution history
- Settings

## Troubleshooting

### "command not found: go"

Make sure Go is installed and in your PATH:

```bash
go version
```

### Build Errors

Clean and rebuild:

```bash
make clean
go mod tidy
make build
```

### Permission Denied

Make scripts executable:

```bash
chmod +x build.sh run.sh
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## Architecture

### Storage Layer

The storage layer uses a simple JSON file database with thread-safe operations:

- `database.go`: Database initialization and operations
- `commands.go`: Command CRUD operations
- `history.go`: History tracking
- `settings.go`: Configuration management

### Command Layer

Each command is implemented as a Cobra command:

- Self-contained functionality
- Clear argument parsing
- Error handling
- User-friendly output

### Utilities

- **colors.go**: Provides color formatting using fatih/color
- **exec.go**: Command execution with both interactive and non-interactive modes
- **table.go**: Table rendering using tablewriter

## Dependencies

- **cobra**: CLI framework
- **fatih/color**: Terminal colors
- **tablewriter**: Table rendering
- **uuid**: UUID generation

## Future Improvements

- [ ] Add unit tests
- [ ] Implement actual marketplace integration
- [ ] Add command tagging and search
- [ ] Support for command templates with variables
- [ ] Shell completion scripts
- [ ] Interactive command builder
- [ ] Command aliases
- [ ] Encrypted secret storage
