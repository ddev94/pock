package commands

import (
	"fmt"
	"os"
	"pock/pkg/pock"

	"github.com/spf13/cobra"
)

// Execute builds and runs the root command.
func Execute() {
	rootCmd := NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// NewRootCommand creates the root CLI command.
func NewRootCommand() *cobra.Command {
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
		Version: pock.Version,
	}

	registerCommands(rootCmd)
	return rootCmd
}
