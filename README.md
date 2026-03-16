<p align="center">
	<img src="assets/icon.svg" alt="Pock icon" width="140" height="140" />
</p>

<h1 align="center">pock</h1>

<p align="center">
	A simple app for saving and reusing terminal commands.
</p>

<p align="center">
	Keep your most-used commands in one place and run them anytime.
</p>

<p align="center">
	<img src="assets/demo.png" alt="Pock demo" width="800" />
</p>

## Overview

`pock` helps you save commands you use again and again, so you do not need to remember or retype them. You can give each command a short name, run it later, and keep your routine tasks organized.

## Highlights

- Save commands with easy-to-remember names
- Run saved commands in seconds
- Keep a history of what you ran with full output logs
- Export and import your command library
- Save script files for later use
- Keep everything on your own computer

## Installation

### Quick Install (macOS/Linux)

Install directly from GitHub releases:

```bash
curl -o- https://raw.githubusercontent.com/ddev94/pock/main/install.sh | bash
```

or with wget:

```bash
wget -qO- https://raw.githubusercontent.com/ddev94/pock/main/install.sh | bash
```

To install a specific version, pass it as an argument:

```bash
curl -o- https://raw.githubusercontent.com/ddev94/pock/main/install.sh | bash -s 1.0.2
```

## Quick Start

You only need a few commands to get started.

### Save something you use often

```bash
pock add hello "echo 'Hello, World!'"
```

This saves the command under the name `hello`.

### Save a script file

```bash
pock add deploy ./deploy.sh -d "Deployment script"
```

When you add a script file, `pock` stores its content so you can run it anytime.

### See what you have saved

```bash
pock list
```

### Run a saved command

```bash
pock run hello
```

## Commands

### `add`

Save a command or script with a short name.

```bash
pock add <name> "<command>" [-d "description"]
pock add <name> ./script.sh [-d "description"]
```

### `list`

Show all saved commands.

```bash
pock list [--stats] [--all]
```

Options:

- `--stats, -s`: Show execution statistics
- `--all, -a`: Show all commands including disabled ones

### `run`

Run a saved command. Output is captured and saved to history.

```bash
pock run <name>
```

### `remove`

Delete a saved command.

```bash
pock remove <name>
```

### `history`

View execution history with optional output viewing.

```bash
pock history [command-name] [--limit 20] [--output]
pock history my-cmd --clear
pock history --clear
```

Options:

- `--limit, -l`: Limit number of entries (default: 20)
- `--output, -o`: Show command output in history
- `--clear`: Clear history (all or for specific command)

### `browse`

Discover and install commands from community-shared libraries.

```bash
pock browse
```

### `register`

Register and publish your commands to share with others.

```bash
pock register
```

### `publish`

Publish your local commands for community use.

```bash
pock publish [--name <command-name>]
```

### `export`

Save your commands to a file for backup or sharing.

```bash
pock export <output-file> [--name <command-name>]
```

### `import`

Import commands from a file or URL.

```bash
pock import <file-or-url> [--force]
```

### `features`

Display upcoming features and development roadmap.

```bash
pock features
```

### `completion`

Generate shell completion scripts for bash, zsh, or fish.

```bash
pock completion bash > /usr/local/etc/bash_completion.d/pock
pock completion zsh > /usr/local/share/zsh/site-functions/_pock
pock completion fish > ~/.config/fish/completions/pock.fish
```

## Example Workflows

### Common shortcuts

```bash
pock add sync-main "git checkout main && git pull --rebase origin main"
pock add deploy "git push origin main"

# Run your commands
pock list
pock run sync-main
```

### Project automation

```bash
pock add dev "npm run dev"
pock add test-all "npm run lint && npm run test && npm run build"
```

### Reusable scripts

```bash
pock add release ./scripts/release.sh -d "Release workflow"
```

### View execution history

```bash
# View all history
pock history

# View history for specific command
pock history dev

# View history with output logs
pock history dev --output

# Clear specific command history
pock history dev --clear
```

## Contributing

We welcome contributions! Here's how you can help:

### Reporting Issues

If you find a bug or have a feature request:

1. Check if the issue already exists in the [GitHub Issues](https://github.com/ddev94/pock/issues)
2. If not, create a new issue with:
   - A clear description of the problem or feature
   - Steps to reproduce (for bugs)
   - Expected vs actual behavior
   - Your environment (OS, Go version, etc.)

### Submitting Changes

1. **Fork the repository**

   ```bash
   git clone https://github.com/ddev94/pock.git
   cd pock
   ```

2. **Create a feature branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add tests if applicable
   - Update documentation as needed

4. **Test your changes**

   ```bash
   make test
   make build
   ```

5. **Commit your changes**

   ```bash
   git commit -m "Add: description of your changes"
   ```

6. **Push to your fork**

   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**
   - Provide a clear description of the changes
   - Reference any related issues
   - Wait for review and address feedback

### Code of Conduct

Be respectful and constructive in all interactions. We aim to maintain a welcoming community for everyone.

## License

ISC
