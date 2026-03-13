# Quick Start Guide

## Installation

### From Source

```bash
# Navigate to the golang directory
cd golang

# Build the binary
go build -o bin/pock ./cmd/pock

# Install globally (optional)
sudo mv bin/pock /usr/local/bin/
```

### Quick Setup with Make

```bash
cd golang
make build    # Build the binary
make install  # Install to /usr/local/bin
```

## Basic Usage

### 1. Add Your First Command

```bash
pock add hello "echo 'Hello, World!'"
```

### 2. List All Commands

```bash
pock list
```

### 3. Run a Command

```bash
pock run hello
```

### 4. Remove a Command

```bash
pock remove hello
```

## Common Workflows

### Daily Commands

```bash
# Add frequently used git commands
pock add gst "git status" -d "Check git status"
pock add glog "git log --oneline --graph --all" -d "Pretty git log"
pock add gp "git push origin main" -d "Push to main"

# Add deployment commands
pock add deploy "npm run build && npm run deploy" -d "Build and deploy"
pock add test "npm test -- --coverage" -d "Run tests with coverage"

# Add docker commands
pock add dps "docker ps -a" -d "List all containers"
pock add dlogs "docker logs -f" -d "Follow docker logs"
```

### Team Collaboration

```bash
# Export your commands
pock export my-commands.json --author "Your Name" --tags "git,docker"

# Share with team (they import)
pock import my-commands.json

# Or import from URL
pock import https://gist.githubusercontent.com/user/id/raw/commands.json
```

### Project-Specific Commands

```bash
# Add project-specific commands
pock add dev "npm run dev" -d "Start development server"
pock add build "npm run build" -d "Build for production"
pock add lint "npm run lint && npm run format" -d "Lint and format code"
pock add db-reset "npm run db:reset && npm run db:seed" -d "Reset and seed database"

# Export for the project
pock export project-commands.json

# Add to .gitignore to keep personal commands private
echo "my-personal-commands.json" >> .gitignore
```

### Viewing History

```bash
# View last 20 executions
pock history

# View last 50 executions
pock history --limit 50

# Clear history
pock history --clear
```

### Configuration

```bash
# Use simple layout (no table)
pock config set listLayout simple

# Use table layout (default)
pock config set listLayout table

# Change date format
pock config set dateFormat iso      # 2024-01-15T10:30:00Z
pock config set dateFormat locale   # 1/15/2024, 10:30:00 AM
pock config set dateFormat relative # 2 hours ago

# View current config
pock config list

# Reset to defaults
pock config reset
```

## Advanced Usage

### Script Files

```bash
# Create a script file
cat > deploy.sh << 'EOF'
#!/bin/bash
echo "Starting deployment..."
npm run build
npm run test
git push origin main
echo "Deployment complete!"
EOF

chmod +x deploy.sh

# Add script as command
pock add deploy "./deploy.sh" -d "Deploy application"

# Run it
pock run deploy
```

### Complex Commands

```bash
# Multi-line commands with && and ||
pock add full-deploy "npm run lint && npm run test && npm run build && npm run deploy" -d "Full CI/CD pipeline"

# Commands with pipes
pock add find-large "find . -type f -size +10M | sort -rh | head -10" -d "Find large files"

# Commands with environment variables
pock add prod-deploy "NODE_ENV=production npm run build && npm run deploy" -d "Production deployment"
```

### Organizing Commands

```bash
# Use descriptive names with prefixes
pock add git-status "git status"
pock add git-push "git push origin main"
pock add git-pull "git pull origin main"

pock add docker-ps "docker ps -a"
pock add docker-stop-all "docker stop \$(docker ps -aq)"
pock add docker-clean "docker system prune -af"

pock add npm-dev "npm run dev"
pock add npm-build "npm run build"
pock add npm-test "npm test"

# List all git commands
pock list | grep git

# View with stats to see most used
pock list --stats
```

## Tips and Tricks

### 1. Quick Aliases

Add these to your shell profile:

```bash
alias h="pock"
alias ha="pock add"
alias hl="pock list"
alias hr="pock run"
alias hh="pock history"
```

### 2. Backup Your Commands

```bash
# Regular backups
pock export ~/Dropbox/pock-backup.json

# Schedule with cron (daily backup)
echo "0 0 * * * pock export ~/backups/pock-\$(date +\%Y\%m\%d).json" | crontab -
```

### 3. Share with Team

```bash
# Create team command library
pock export team-commands.json --author "Team Name"

# Upload to GitHub Gist or company wiki
# Team members can then import
```

### 4. Command Templates

```bash
# Add template commands with placeholders
pock add deploy-branch "git checkout \$BRANCH && git push origin \$BRANCH" -d "Deploy specific branch"

# Note: You'll need to edit the command when running or create specific versions
pock add deploy-dev "git checkout dev && git push origin dev"
pock add deploy-staging "git checkout staging && git push origin staging"
```

### 5. Database Location

```bash
# Find your database
ls -la ~/.local/share/pock/db.json

# Backup database directly
cp ~/.local/share/pock/db.json ~/pock-backup-$(date +%Y%m%d).json

# View database contents
cat ~/.local/share/pock/db.json | jq .
```

## Troubleshooting

### Command Not Found

```bash
# Make sure pock is in PATH
which pock

# If not, add to PATH in your shell profile
export PATH=$PATH:/path/to/pock/bin
```

### Permission Denied

```bash
# Make sure binary is executable
chmod +x /usr/local/bin/pock
```

### Commands Not Working

```bash
# Check if command exists
pock list

# View execution history for errors
pock history

# Try running command directly in terminal first
echo "test command"
```

## Getting Help

```bash
# General help
pock --help

# Command-specific help
pock add --help
pock list --help
pock run --help

# Version info
pock --version
```

## Next Steps

1. Read the full [README.md](README.md) for detailed documentation
2. Check [DEVELOPMENT.md](DEVELOPMENT.md) for development guide
3. See [COMPARISON.md](COMPARISON.md) for TypeScript vs Go comparison
4. Start adding your frequently used commands!

## Examples by Use Case

### DevOps Engineer

```bash
pock add k8s-pods "kubectl get pods -A"
pock add k8s-logs "kubectl logs -f"
pock add terraform-plan "terraform plan -out=plan.tfplan"
pock add ansible-deploy "ansible-playbook -i inventory deploy.yml"
```

### Web Developer

```bash
pock add dev "npm run dev"
pock add build "npm run build"
pock add test "npm test"
pock add lint "eslint . && prettier --check ."
pock add db "docker-compose up -d postgres"
```

### Data Scientist

```bash
pock add jupyter "jupyter lab --port=8888"
pock add train "python train.py --epochs=100"
pock add tensorboard "tensorboard --logdir=./logs"
pock add notebook "jupyter notebook --no-browser"
```

### System Administrator

```bash
pock add disk-usage "df -h | grep -v tmpfs"
pock add memory "free -h"
pock add processes "ps aux | head -20"
pock add network "netstat -tuln"
pock add logs "tail -f /var/log/syslog"
```

Enjoy using pock! 🚀
