package cmd

import "github.com/spf13/cobra"

func registerCommands(rootCmd *cobra.Command) {
	// Core commands - always enabled
	rootCmd.AddCommand(
		NewAddCommand(),
		NewListCommand(),
		NewRunCommand(),
		NewRemoveCommand(),
	)

	// Optional commands - controlled by feature flags
	if EnableHistoryCommand {
		rootCmd.AddCommand(NewHistoryCommand())
	}
	if EnableExportCommand {
		rootCmd.AddCommand(NewExportCommand())
	}
	if EnableImportCommand {
		rootCmd.AddCommand(NewImportCommand())
	}
	if EnableInstallCommand {
		rootCmd.AddCommand(NewInstallCommand())
	}
	if EnableBrowseCommand {
		rootCmd.AddCommand(NewBrowseCommand())
	}
	if EnablePublishCommand {
		rootCmd.AddCommand(NewPublishCommand())
	}
}
