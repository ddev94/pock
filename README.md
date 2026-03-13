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

## Overview

`pock` helps you save commands you use again and again, so you do not need to remember or retype them. You can give each command a short name, run it later, and keep your routine tasks organized.

## Highlights

- Save commands with easy-to-remember names
- Run saved commands in seconds
- Keep a history of what you ran
- Export and import your command library
- Save script files for later use
- Keep everything on your own computer

## Installation

### Install from the macOS package

1. Download the latest `.pkg` installer.
2. Open the installer file.
3. Follow the installation steps on screen.
4. Open Terminal and run:

```bash
pock --help
```

If you see the help message, `pock` is installed correctly.

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

When you add a script file, `pock` keeps its own copy so you can still use it later.

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
pock list [--stats]
```

### `run`

Run a saved command.

```bash
pock run <name>
```

### `remove`

Delete a saved command you no longer need.

```bash
pock remove <name>
```

### `history`

See what you ran before.

```bash
pock history [--limit 20]
pock history --clear
```

### `export`

Save your commands to a file for backup or sharing.

```bash
pock export <output-file> [--name <command-name>]
```

### `import`

Bring commands in from a file or link.

```bash
pock import <file-or-url> [--force]
```

### `config`

Change app settings.

```bash
pock config set <key> <value>
pock config get <key>
pock config list
```

## Example Workflows

### Common shortcuts

```bash
pock add sync-main "git checkout main && git pull --rebase origin main"
pock add publish "git push origin main"
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

## Documentation

More detailed guides are available here:

- [QUICKSTART.md](QUICKSTART.md)
- [DEVELOPMENT.md](DEVELOPMENT.md)
- [ARCHITECTURE.md](ARCHITECTURE.md)
- [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)
- [COMPARISON.md](COMPARISON.md)

## License

ISC
