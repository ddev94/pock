# pock - Go Implementation

A powerful command-line tool for saving, managing, and sharing your frequently used commands - now written in Go!

## Features

- 💾 Save and manage frequently used commands
- ⚡ Quick command execution with simple aliases
- 📝 Command history tracking
- 🌐 Marketplace for sharing and discovering commands
- 📤 Export commands to shareable JSON files
- 📥 Import commands from files or URLs
- 🔍 Search and browse community commands
- 📦 Simple and intuitive command structure
- 🚀 Written in Go for performance and cross-platform compatibility

## Installation

### From Source

```bash
cd golang
go build -o pock ./cmd/pock
sudo mv pock /usr/local/bin/
```

### Development

```bash
cd golang
go run ./cmd/pock <command>
```

## Usage

### Basic Command Structure

```bash
pock <command> [options]
```

### Available Commands

#### Add Command

Save a new command for later use:

```bash
pock add <name> "<command>" [-d "description"]
```

Example:

```bash
pock add hello "echo 'Hello, World!'" -d "A simple hello world command"
pock add deploy "git push origin main && npm run deploy" -d "Deploy to production"
```

#### List Commands

View all saved commands:

```bash
pock list [--stats]
```

Options:

- `--stats, -s` - Show execution statistics for each command

#### Run Command

Execute a saved command:

```bash
pock run <name>
```

#### Remove Command

Delete a saved command:

```bash
pock remove <name>
```

#### History

View command execution history:

```bash
pock history [--limit 20]
```

#### Export

Export commands to a JSON file:

```bash
pock export <output-file> [--name <command-name>]
```

#### Import

Import commands from a JSON file or URL:

```bash
pock import <file-or-url> [--force]
```

#### Config

Manage configuration settings:

```bash
pock config set <key> <value>
pock config get <key>
pock config list
```

## Project Structure

```
golang/
├── cmd/
│   └── pock/
│       └── main.go          # Entry point
├── internal/
│   ├── commands/
│   │   ├── add.go           # Add command
│   │   ├── list.go          # List command
│   │   ├── run.go           # Run command
│   │   ├── remove.go        # Remove command
│   │   ├── history.go       # History command
│   │   ├── export.go        # Export command
│   │   ├── import.go        # Import command
│   │   ├── config.go        # Config command
│   │   ├── install.go       # Install command
│   │   ├── browse.go        # Browse command
│   │   └── publish.go       # Publish command
│   ├── storage/
│   │   ├── types.go         # Data types
│   │   ├── database.go      # Database operations
│   │   ├── commands.go      # Command storage
│   │   ├── history.go       # History storage
│   │   └── settings.go      # Settings storage
│   └── utils/
│       ├── colors.go        # Color utilities
│       ├── exec.go          # Command execution
│       └── table.go         # Table rendering
├── go.mod
└── README.md
```

## License

ISC
