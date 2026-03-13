package main

import (
	"fmt"
	"os"
	"pock/internal/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pock",
		Short: "pock - Command manager for saving, sharing, and running commands",
		Long: `A powerful command-line tool for saving, managing, and sharing your frequently used commands.

Features:
  • Save and manage frequently used commands
  • Quick command execution with simple aliases
  • Command history tracking
  • Marketplace for sharing and discovering commands
  • Export/Import commands to shareable JSON files
  • Search and browse community commands`,
		Version: "1.0.0",
	}

	// Add all commands
	rootCmd.AddCommand(commands.NewAddCommand())
	rootCmd.AddCommand(commands.NewListCommand())
	rootCmd.AddCommand(commands.NewRunCommand())
	rootCmd.AddCommand(commands.NewRemoveCommand())
	rootCmd.AddCommand(commands.NewHistoryCommand())
	rootCmd.AddCommand(commands.NewExportCommand())
	rootCmd.AddCommand(commands.NewImportCommand())
	rootCmd.AddCommand(commands.NewConfigCommand())
	rootCmd.AddCommand(commands.NewInstallCommand())
	rootCmd.AddCommand(commands.NewBrowseCommand())
	rootCmd.AddCommand(commands.NewPublishCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
