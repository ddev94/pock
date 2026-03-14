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

## Getting Started

### 1. Download Dependencies

```bash
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

### Build macOS Package

```bash
make package
```

This creates a `.pkg` installer for macOS distribution.

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

# Add a script file
pock add release ./scripts/release.sh -d "Release workflow"
```

### Listing Commands

```bash
# List all enabled commands
pock list

# Use alias
pock ls

# List with execution statistics
pock list --stats

# List all commands including disabled ones
pock list --all
```

### Running Commands

```bash
# Run a saved command (output is captured and saved)
pock run hello

# The command runs with animated spinner during lookup
```

### Managing History

```bash
# View execution history
pock history

# View last 10 entries
pock history --limit 10

# View history for specific command
pock history deploy

# View history with output logs
pock history deploy --output

# Clear all history
pock history --clear

# Clear history for specific command
pock history deploy --clear
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

### Command Aliases

Short aliases for faster workflow:

```bash
pock ls              # Same as: pock list
pock rm my-cmd       # Same as: pock remove my-cmd
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
- `commands.go`: Command CRUD operations with enable/disable support
- `history.go`: History tracking with command output logs
- `settings.go`: Configuration management
- `types.go`: Type definitions for all data structures

### Command Layer

Each command is implemented as a Cobra command:

- Self-contained functionality
- Clear argument parsing
- Error handling
- User-friendly output with colors

### Feature Flags

Commands can be enabled/disabled via constants in `internal/commands/features.go`:

```go
const (
    EnableBrowseCommand  = false // Marketplace integration
    EnableInstallCommand = false // Not yet implemented
    EnablePublishCommand = false // Not yet implemented
    EnableExportCommand  = true  // Enabled
    EnableImportCommand  = true  // Enabled
    EnableHistoryCommand = true  // Enabled
)
```

### UI Components

- **colors.go**: Provides color formatting using Charm Bracelet's Lipgloss
- **exec.go**: Command execution with both interactive and non-interactive modes, captures output
- **table.go**: Responsive table rendering using Lipgloss, auto-adjusts to terminal width
- **run.go**: Implements Bubble Tea spinner for command lookup animation

## Dependencies

### Core Dependencies

- **cobra**: CLI framework for command structure
- **lipgloss**: Terminal styling and colors
- **bubbletea**: Terminal UI framework for interactive components
- **bubbles**: Pre-built TUI components (spinner, table)
- **uuid**: UUID generation for entities

### Additional Dependencies

- **charmbracelet/x/ansi**: ANSI escape code handling
- **golang.org/x/term**: Terminal width detection for responsive tables

All dependencies are from Charm Bracelet ecosystem for consistent, beautiful terminal UI.

## Feature Development

### Adding a New Command

1. Create command file in `internal/commands/`:

   ```go
   // my_command.go
   func NewMyCommand() *cobra.Command {
       cmd := &cobra.Command{
           Use:   "my-command",
           Short: "Description",
           RunE: func(cmd *cobra.Command, args []string) error {
               // Implementation
               return nil
           },
       }
       return cmd
   }
   ```

2. Add feature flag in `internal/commands/features.go`:

   ```go
   const EnableMyCommand = true
   ```

3. Register in `internal/commands/register.go`:
   ```go
   if EnableMyCommand {
       rootCmd.AddCommand(NewMyCommand())
   }
   ```

### Modifying Table Rendering

Table widths are configured in `internal/utils/table.go`:

```go
idealWidths := map[string]int{
    "Name":        20,
    "Command":     50,
    "Description": 30,
}
maxWidths := map[string]int{
    "Name":        30,
    "Command":     80,
    "Description": 50,
}
```

Tables automatically shrink proportionally when terminal width is too narrow.

## Future Improvements

- [x] Command aliases (ls, rm)
- [x] Command execution output capture
- [x] History viewing with output logs
- [x] Responsive table rendering
- [x] Bubble Tea spinner for command lookup
- [x] Feature flags for commands
- [x] Command-specific history viewing and clearing
- [ ] Add comprehensive unit tests
- [ ] Implement actual marketplace integration
- [ ] Add command tagging and search
- [ ] Support for command templates with variables
- [ ] Interactive command builder with Bubble Tea
- [ ] Encrypted secret storage for sensitive commands
